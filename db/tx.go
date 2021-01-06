package db

import (
	"fmt"

	"github.com/BOPR/bazooka"
	"github.com/BOPR/config"
	"github.com/BOPR/core"
	"github.com/BOPR/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/kilic/bn254/bls"
)

// InsertTx tx into the DB
func (DBI *DB) InsertTx(tx *core.Tx) error {
	// if tx is a create2transfer tx add it to the relayer pool
	if tx.Type == core.TX_CREATE_2_TRANSFER {
		return DBI.InsertRelayPacket(tx.Data, tx.Signature)
	}
	return DBI.Instance.Create(tx).Error
}

// PopTxs
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
	// select N number of transactions which are pending in mempool and
	if err := tx.Limit(config.GlobalCfg.TxsPerBatch).Where(&core.Tx{Status: core.TX_STATUS_PENDING, Type: txType}).Find(&pendingTxs).Error; err != nil {
		DBI.Logger.Error("error while fetching pending transactions", err)
		tx.Rollback()
		return txs, err
	}

	DBI.Logger.Info("Found txs", "pendingTxs", len(pendingTxs))
	var ids []string
	for _, tx := range pendingTxs {
		ids = append(ids, tx.ID)
	}

	// update the transactions from pending to processing
	errs := tx.Table("txes").Where("id IN (?)", ids).Updates(map[string]interface{}{"status": core.TX_STATUS_PROCESSING}).GetErrors()
	if err != nil {
		tx.Rollback()
		DBI.Logger.Error("errors while processing transactions", errs)
		return
	}
	return pendingTxs, tx.Commit().Error
}

func (DBI *DB) FetchTxType() (txType uint64, err error) {
	// find out which txType has the highest count
	var maxTxType uint64
	var maxCount uint64
	txTypes := []uint64{core.TX_TRANSFER_TYPE}
	for _, txType := range txTypes {
		count, err := DBI.GetCountPerTxType(txType)
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

func (DBI *DB) GetCountPerTxType(txType uint64) (uint64, error) {
	var count uint64
	err := DBI.Instance.Model(&core.Tx{}).Where("type = ? AND status = ?", txType, core.TX_STATUS_PENDING).Count(&count).Error
	return count, err
}

func (DBI *DB) GetTx() (tx []core.Tx, err error) {
	err = DBI.Instance.First(&tx).Error
	if err != nil {
		return tx, err
	}
	return
}

func (DBI *DB) FetchMPWithID(id uint64, stateMP *bazooka.StateMerkleProof) (err error) {
	leaf, err := DBI.GetStateByIndex(id)
	if err != nil {
		fmt.Println("error while getting leaf", err)
		return
	}
	siblings, err := DBI.GetSiblings(leaf.Path)
	if err != nil {
		fmt.Println("error while getting siblings", err)
		return
	}
	accMP := bazooka.NewStateMerkleProof(leaf, siblings)
	*stateMP = accMP
	return nil
}

func GetWitnessTranfer(bz bazooka.Bazooka, DBI DB, tx core.Tx) (fromMerkleProof, toMerkleProof bazooka.StateMerkleProof, txDBConn DB, err error) {
	dbCopy, _ := NewDB()

	// fetch from state MP
	err = DBI.FetchMPWithID(tx.From, &fromMerkleProof)
	if err != nil {
		return
	}
	toState, err := DBI.GetStateByIndex(tx.To)
	if err != nil {
		return
	}
	var toSiblings []core.UserState

	newFrom, newTo, err := bazooka.ApplyTx(bz, fromMerkleProof.State.Data, toState.Data, tx)
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

// GetVerificationData fetches all the data required to prove validity fo transaction
func GetVerificationData(bz bazooka.Bazooka, DBI *DB, tx *core.Tx) (fromMerkleProof, toMerkleProof bazooka.StateMerkleProof, txDBConn DB, err error) {
	switch txType := tx.Type; txType {
	case core.TX_TRANSFER_TYPE:
		return GetWitnessTranfer(bz, *DBI, *tx)
	case core.TX_CREATE_2_TRANSFER:
		return GetWitnessTranfer(bz, *DBI, *tx)
	case core.TX_MASS_MIGRATIONS:
		return GetWitnessTranfer(bz, *DBI, *tx)
	default:
		fmt.Println("TxType didnt match any options", tx.Type)
		return
	}
}

// Validate creates proofs for validating txs and returns new root post validation
func Validate(bz *bazooka.Bazooka, DBI *DB, currentRoot core.ByteArray, tx *core.Tx, isSyncing bool) (newRoot core.ByteArray, err error) {
	fromStateProof, toStateProof, txDBConn, err := GetVerificationData(*bz, DBI, tx)
	if err != nil {
		return
	}

	newRoot, err = bazooka.ProcessTx(*bz, currentRoot, *tx, fromStateProof, toStateProof)
	if err != nil {
		if txDBConn.Instance != nil {
			txDBConn.Instance.Rollback()
			txDBConn.Close()
		}
		return
	}

	if !isSyncing {
		err = authenticate(bz, DBI, tx)
		if err != nil {
			txDBConn.Instance.Rollback()
			txDBConn.Close()
			return
		}
	}

	if txDBConn.Instance != nil {
		txDBConn.Instance.Commit()
		txDBConn.Close()
	}

	return
}

func authenticate(bz *bazooka.Bazooka, DBI *DB, tx *core.Tx) error {
	fromState, err := DBI.GetStateByIndex(tx.From)
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

	err = authenticateTx(bz, DBI, *tx, fromAcc.PublicKey)
	if err != nil {
		return err
	}

	return nil
}

// ProcessTxs processes all trasnactions and returns the commitment list
func ProcessTxs(bz *bazooka.Bazooka, DBI *DB, txs []core.Tx, isSyncing bool) (commitments []core.Commitment, err error) {
	if len(txs) == 0 {
		return commitments, core.ErrNoTxsFound
	}
	for i, tx := range txs {
		rootAcc, err := DBI.GetRoot()
		if err != nil {
			return commitments, err
		}
		currentRoot, err := core.HexToByteArray(rootAcc.Hash)
		if err != nil {
			return commitments, err
		}
		newRoot, err := Validate(bz, DBI, currentRoot, &tx, isSyncing)
		if err != nil {
			return commitments, err
		}
		if i%int(config.GlobalCfg.TxsPerBatch) == 0 {
			txInCommitment := txs[i : i+int(config.GlobalCfg.TxsPerBatch)]
			aggregatedSig, err := aggregateSignatures(txInCommitment)
			if err != nil {
				if isSyncing && err == core.ErrSignatureNotPresent {
					continue
				} else {
					return commitments, err
				}
			}
			commitment := core.NewCommitment(0, 0, txInCommitment, tx.Type, newRoot, core.ByteArray{}, aggregatedSig.ToBytes())
			commitments = append(commitments, commitment)
		}
	}

	return commitments, nil
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

func authenticateTx(b *bazooka.Bazooka, IDB *DB, tx core.Tx, pubkeySender []byte) error {
	opts := bind.CallOpts{From: config.OperatorAddress}
	solPubkeySender, err := core.Pubkey(pubkeySender).ToSol()
	if err != nil {
		return err
	}
	signature, err := core.BytesToSolSignature(tx.Signature)
	if err != nil {
		return err
	}

	switch tx.Type {
	case core.TX_TRANSFER_TYPE:
		err = b.SC.Transfer.Validate(&opts, tx.Data, signature, solPubkeySender, wallet.DefaultDomain)
		if err != nil {
			return err
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
		err = b.SC.Create2Transfer.Validate(&opts, tx.Data, signature, solPubkeySender, solPubkeyReceiver, wallet.DefaultDomain)
		if err != nil {
			return err
		}
	case core.TX_MASS_MIGRATIONS:
		err = b.SC.MassMigration.Validate(&opts, tx.Data, signature, solPubkeySender, wallet.DefaultDomain)
		if err != nil {
			return err
		}
	}
	return nil
}
