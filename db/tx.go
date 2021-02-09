package db

import (
	"errors"
	"fmt"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/core"
	"github.com/BOPR/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	"github.com/kilic/bn254/bls"
)

var (
	// ErrBadSignature error bad signature
	ErrBadSignature = errors.New("Invalid signature")
)

// InsertTx inserts a transaction into the DB
func (DBI *DB) InsertTx(tx *core.Tx) error {
	// if tx is a create2transfer tx add it to the relayer pool
	if tx.Type == core.TX_CREATE_2_TRANSFER {
		return DBI.InsertRelayPacket(tx.Data, tx.Signature)
	}

	return DBI.Instance.Create(tx).Error
}

// GetTxByHash fetches transaction by hash
func (DBI *DB) GetTxByHash(hash string) (*core.Tx, error) {
	var tx core.Tx
	if err := DBI.Instance.Model(&tx).Scopes(QueryByTxHash(hash)).First(&tx).Error; err != nil {
		return &tx, err
	}
	return &tx, nil
}

// GetPendingNonce fetches the pending nonce for a particular state
// to be used by client to determine tx nonce
func (DBI *DB) GetPendingNonce(senderStateID uint64) (nonce uint64, err error) {
	// fetch nonce in state
	state, err := DBI.GetStateByIndex(senderStateID)
	if err != nil {
		return
	}

	_, _, stateNonce, _, err := DBI.Bazooka.DecodeState(state.Data)
	if err != nil {
		return
	}

	// fetch all txs by a sender with status pending sorted by nonce
	var latestTx core.Tx
	err = DBI.Instance.Where(&core.Tx{Status: core.TX_STATUS_PENDING, From: senderStateID}).Order("nonce desc").First(&latestTx).Error
	if gorm.IsRecordNotFoundError(err) {
		return stateNonce.Uint64(), nil
	}

	if err != nil {
		return 0, err
	}

	if stateNonce.Uint64() > latestTx.Nonce {
		return stateNonce.Uint64(), nil
	}

	return latestTx.Nonce, nil
}

// PopTxs pops tranasctions from the tx pool
func (DBI *DB) PopTxs() (txs []core.Tx, err error) {
	txType, err := DBI.FetchTxType()
	tx := DBI.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return txs, err
	}

	var pendingTxs []core.Tx
	totalTxs := DBI.Cfg.TxsPerCommitment * DBI.Cfg.MaxCommitmentsPerBatch

	// select N number of transactions which are pending in mempool and
	if err := tx.Limit(totalTxs).Where(&core.Tx{Status: core.TX_STATUS_PENDING, Type: txType}).Find(&pendingTxs).Error; err != nil {
		DBI.Logger.Error("error while fetching pending transactions", err)
		tx.Rollback()
		return txs, err
	}

	DBI.Logger.Info("Found pending txs", "txCount", len(pendingTxs))

	// update the status fo pending transactions to processing
	err = DBI.UpdateTxStatuses(pendingTxs, core.TX_STATUS_PROCESSING)
	if err != nil {
		tx.Rollback()
		return txs, err
	}

	return pendingTxs, tx.Commit().Error
}

// UpdateTxStatuses updates the tx status
func (DBI *DB) UpdateTxStatuses(txList []core.Tx, status int) error {
	var ids []string
	for _, tx := range txList {
		ids = append(ids, tx.ID)
	}
	err := DBI.Instance.Table("txes").Where("id IN (?)", ids).Updates(map[string]interface{}{"status": status}).Error
	if err != nil {
		DBI.Logger.Error("errors while processing transactions", "error", err)
		return err
	}

	return nil
}

// FetchTxType finds out transactions with highest count in the DB
func (DBI *DB) FetchTxType() (txType uint64, err error) {
	// find out which txType has the highest count
	var maxTxType uint64
	var maxCount uint64
	txTypes := []uint64{core.TX_TRANSFER_TYPE}
	for _, txType := range txTypes {
		count, err := DBI.getCountPerTxType(txType)
		if err != nil {
			return 0, err
		}
		if count > maxCount {
			maxTxType = txType
			maxCount = count
		}
	}
	return maxTxType, nil
}

// getWitnessTranfer fetches the witness for transfer transaction and create2transfer
// amongst other thins it also returns a DB transaction handler that can be used to rollback changes made to state tree while creating witness
func getWitnessTranfer(bz bazooka.Bazooka, DBI DB, tx core.Tx) (fromMerkleProof, toMerkleProof bazooka.StateMerkleProof, txDBConn DB, err error) {
	dbCopy, _ := NewDB(bz.Cfg)

	from, to, err := bz.FetchFromAndToStateIDs(tx)
	if err != nil {
		return
	}

	// fetch from state MP
	err = DBI.fetchMPWithID(from, &fromMerkleProof)
	if err != nil {
		return
	}
	toState, err := DBI.GetStateByIndex(to)
	if err != nil {
		return
	}
	var toSiblings []core.UserState

	newFrom, newTo, err := bz.ApplyTx(fromMerkleProof.State.Data, toState.Data, tx)
	if err != nil {
		return
	}

	mysqlTx := dbCopy.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			mysqlTx.Rollback()
		}
	}()
	dbCopy.Instance = mysqlTx

	// apply the new from leaf
	currentFromStateCopy := fromMerkleProof.State
	currentFromStateCopy.Data = newFrom
	err = dbCopy.UpdateState(currentFromStateCopy)
	if err != nil {
		return
	}

	// create witness for to leaf
	toSiblings, err = dbCopy.GetSiblings(toState.Path)
	if err != nil {
		return
	}
	toMerkleProof = bazooka.NewStateMerkleProof(toState, toSiblings)

	// apply the new to leaf
	toState.Data = newTo
	err = dbCopy.UpdateState(toState)
	if err != nil {
		return
	}

	return fromMerkleProof, toMerkleProof, dbCopy, nil
}

// getWitnessMassMigration fetches the witness for transfer mass migrations
// amongst other thins it also returns a DB transaction handler that can be used to rollback changes made to state tree while creating witness
func getWitnessMassMigration(bz bazooka.Bazooka, DBI DB, tx core.Tx) (fromMerkleProof, toMerkleProof bazooka.StateMerkleProof, txDBConn DB, err error) {
	dbCopy, _ := NewDB(bz.Cfg)

	from, _, err := bz.FetchFromAndToStateIDs(tx)
	if err != nil {
		return
	}

	// fetch from state MP
	err = DBI.fetchMPWithID(from, &fromMerkleProof)
	if err != nil {
		return
	}

	newFrom, _, err := bz.ApplyTx(fromMerkleProof.State.Data, []byte(""), tx)
	if err != nil {
		return
	}

	mysqlTx := dbCopy.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			mysqlTx.Rollback()
		}
	}()
	dbCopy.Instance = mysqlTx

	// apply the new from leaf
	currentFromStateCopy := fromMerkleProof.State
	currentFromStateCopy.Data = newFrom
	err = dbCopy.UpdateState(currentFromStateCopy)
	if err != nil {
		return
	}

	return fromMerkleProof, toMerkleProof, dbCopy, nil
}

// GetVerificationDataAndApply fetches all the proofs required to prove validity for a transaction and also applies the transaction
// on the respective states
func GetVerificationDataAndApply(bz bazooka.Bazooka, DBI *DB, tx *core.Tx) (fromMerkleProof, toMerkleProof bazooka.StateMerkleProof, txDBConn DB, err error) {
	switch txType := tx.Type; txType {
	case core.TX_TRANSFER_TYPE:
		return getWitnessTranfer(bz, *DBI, *tx)
	case core.TX_CREATE_2_TRANSFER:
		return getWitnessTranfer(bz, *DBI, *tx)
	case core.TX_MASS_MIGRATIONS:
		return getWitnessMassMigration(bz, *DBI, *tx)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return
	}
}

// ValidateAndApplyTx validates and applies transaction by calling on-chain functions
// It checks for state validity as well as account validity
// It returns the newRoot of state tree post execution of transaction
func ValidateAndApplyTx(bz *bazooka.Bazooka, DBI *DB, currentRoot core.ByteArray, tx *core.Tx, isSyncing bool) (newRoot core.ByteArray, err error) {
	// fetches all required proofs for the transaction
	fromStateProof, toStateProof, txDBConn, err := GetVerificationDataAndApply(*bz, DBI, tx)
	if err != nil {
		return
	}

	DBI.Logger.Debug("Fetched all verification for transaction", "txType", tx.Type)

	// calls on-chain function that validates transaction
	newRoot, err = bz.ProcessTx(currentRoot, *tx, fromStateProof, toStateProof)

	// if transaction is declared to be invalid we rollback the state updates made during merkle proof creation
	if err != nil {
		DBI.Logger.Debug("State validation of transaction complete", "status", "fail", "tx", tx.TxHash, "error", err)
		if txDBConn.Instance != nil {
			txDBConn.Instance.Rollback()
			txDBConn.Close()
		}
		return
	}

	DBI.Logger.Debug("State validation of transaction complete", "status", "success", "tx", tx.TxHash, "newRoot", newRoot)

	// if we arent syncing, we need to authenticate the batch
	// i.e we need to check signatures in the transactions
	// if we are syncing we dont have signatures on-chain, we have only aggregated signature
	// so we skip individual signature check
	if !isSyncing {
		err = authenticate(bz, DBI, tx)
		if err != nil {
			txDBConn.Instance.Rollback()
			txDBConn.Close()
			return
		}
		DBI.Logger.Debug("Tx successfully authenticated", "tx", tx.TxHash)
	}

	// if all goes well commit the DB transaction
	if txDBConn.Instance != nil {
		txDBConn.Instance.Commit()
		txDBConn.Close()
	}

	return
}

// ProcessTxs processes transactions using the on-chain contracts and creates a commitment list
// It returns the list of commitments for all valid transactions
func ProcessTxs(bz *bazooka.Bazooka, DBI *DB, txs []core.Tx, txsPerCommitment []int, isSyncing bool) (commitments []core.Commitment, err error) {
	// error out if we dont have commitments to process
	if len(txs) == 0 {
		return commitments, core.ErrNoTxsFound
	}

	// tracks commitment under formation
	currentCommitmentIdx := 0

	var processedTxs []core.Tx
	var revertedTxs []core.Tx

	for i, tx := range txs {
		// fetch the current root
		currentRoot, err := DBI.GetStateRootHash()
		if err != nil {
			return commitments, err
		}

		// validate and apply the transaction
		newRoot, err := ValidateAndApplyTx(bz, DBI, currentRoot, &tx, isSyncing)

		// if transaction validation errors out, add it to reverted txs list
		// and skip rest of the loop
		if err != nil {
			DBI.Logger.Info("Reverting tx", "hash", tx.TxHash, "error", err)
			revertedTxs = append(revertedTxs, tx)
			continue
		}

		processedTxs = append(processedTxs, tx)

		// if num of transactions that are to be packed in the current
		// commitment is reached create a commitment
		// NOTE: If a transaction in commitment reverts that commitment will have smaller size
		if i%txsPerCommitment[currentCommitmentIdx] == 0 {
			// create a new commitment
			var commitment core.Commitment

			// pick all transactions executed successfully so far
			// and empty out the processed txs list for next commitment
			txInCommitment := processedTxs
			processedTxs = nil

			// aggregate all signatures in the transactions
			aggregatedSig, err := aggregateSignatures(txInCommitment)

			// if we are syncing and we dont have a signature
			// create a commitment without it
			if isSyncing && err == core.ErrSignatureNotPresent {
				commitment = core.NewCommitment(txInCommitment, tx.Type, newRoot, []byte(""))
			} else if err != nil {
				return commitments, err
			} else {
				commitment = core.NewCommitment(txInCommitment, tx.Type, newRoot, aggregatedSig.ToBytes())
			}

			// append to list of commitments
			commitments = append(commitments, commitment)

			// doesnt increment currentCommitmentIdx if this is the first commit
			if len(commitments) != 1 {
				currentCommitmentIdx++
			}
		}
	}

	// update the statues for processed and reverted txs
	// we dont bubble up the errors incase we arent able to update tx status
	// we log the error and move on
	err = DBI.UpdateTxStatuses(processedTxs, core.TX_STATUS_PROCESSED)
	if err != nil {
		DBI.Logger.Error("Unable to update tx status for processed txs", "count", len(processedTxs))
		return commitments, nil
	}
	err = DBI.UpdateTxStatuses(revertedTxs, core.TX_STATUS_REVERTED)
	if err != nil {
		DBI.Logger.Error("Unable to update tx status for reverted txs", "count", len(revertedTxs))
		return commitments, nil
	}

	return commitments, nil
}

// checks transaction signature
func authenticate(bz *bazooka.Bazooka, DBI *DB, tx *core.Tx) error {
	from, _, err := bz.FetchFromAndToStateIDs(*tx)
	if err != nil {
		return err
	}
	fromState, err := DBI.GetStateByIndex(from)
	if err != nil {
		return err
	}

	accID, _, _, _, err := bz.DecodeState(fromState.Data)
	if err != nil {
		return err
	}

	params, err := DBI.GetParams()
	if err != nil {
		return err
	}

	path, err := core.SolidityPathToNodePath(accID.Uint64(), params.MaxDepth)
	if err != nil {
		return err
	}

	fromAcc, err := DBI.GetAccountLeafByPath(path)
	if err != nil {
		return err
	}

	err = checkSignature(bz, DBI, *tx, fromAcc.PublicKey)
	if err != nil {
		return err
	}
	return nil
}

// checkSignature calls the on-chain contract to verify signature in the transaction
// returns an error is the signature is invalid
func checkSignature(b *bazooka.Bazooka, IDB *DB, tx core.Tx, pubkeySender []byte) error {
	opts := bind.CallOpts{From: common.HexToAddress(b.Cfg.OperatorAddress)}

	solPubkeySender, err := core.Pubkey(pubkeySender).ToSol()
	if err != nil {
		return err
	}

	// converts signature bytes to SolSignature
	signature, err := core.BytesToSolSignature(tx.Signature)
	if err != nil {
		return err
	}

	switch tx.Type {
	case core.TX_TRANSFER_TYPE:
		ok, err := b.SC.Transfer.Validate(&opts, tx.Data, signature, solPubkeySender, core.StringToByteArray(b.Cfg.AppID))
		if err != nil {
			return err
		}
		if !ok {
			return ErrBadSignature
		}
	case core.TX_CREATE_2_TRANSFER:
		_, _, toAccID, _, _, _, _, err := b.DecodeCreate2Transfer(tx.Data)
		if err != nil {
			return err
		}
		acc, err := IDB.GetAccountLeafByID(toAccID.Uint64())
		if err != nil {
			return err
		}
		solPubkeyReceiver, err := core.Pubkey(acc.PublicKey).ToSol()
		if err != nil {
			return err
		}
		ok, err := b.SC.Create2Transfer.Validate(&opts, tx.Data, signature, solPubkeySender, solPubkeyReceiver, core.StringToByteArray(b.Cfg.AppID))
		if err != nil {
			return err
		}
		if !ok {
			return ErrBadSignature
		}
	case core.TX_MASS_MIGRATIONS:
		ok, err := b.SC.MassMigration.Validate(&opts, tx.Data, signature, solPubkeySender, core.StringToByteArray(b.Cfg.AppID))
		if err != nil {
			return err
		}
		if !ok {
			return ErrBadSignature
		}
	}
	return nil
}

// generates aggregated signature for commitment
func aggregateSignatures(txs []core.Tx) (aggregatedSig bls.Signature, err error) {
	var signatures []*bls.Signature
	for _, tx := range txs {
		if tx.Signature == nil {
			return aggregatedSig, core.ErrSignatureNotPresent
		}
		sig, err := wallet.BytesToSignature(tx.Signature)
		if err != nil {
			return aggregatedSig, err
		}
		signatures = append(signatures, &sig)
	}
	return wallet.NewAggregateSignature(signatures)
}

// fetches state merkle proof using state ID
// it returns only error but populates the stateMerkleProof
func (DBI *DB) fetchMPWithID(id uint64, stateMP *bazooka.StateMerkleProof) (err error) {
	// fetch leaf
	leaf, err := DBI.GetStateByIndex(id)
	if err != nil {
		DBI.Logger.Error("error while getting state", "state-id", id)
		return
	}

	// fetch siblings for leaf
	siblings, err := DBI.GetSiblings(leaf.Path)
	if err != nil {
		fmt.Println("error while getting siblings", err)
		return
	}

	MP := bazooka.NewStateMerkleProof(leaf, siblings)
	*stateMP = MP

	// if all good return no error
	return nil
}

// fetches the tx count per txType from the pending pool
func (DBI *DB) getCountPerTxType(txType uint64) (uint64, error) {
	var count uint64
	err := DBI.Instance.Model(&core.Tx{}).Where("type = ? AND status = ?", txType, core.TX_STATUS_PENDING).Count(&count).Error
	return count, err
}
