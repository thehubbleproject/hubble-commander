package bazooka

import (
	"math/big"
	"strings"

	"github.com/BOPR/contracts/accountregistry"
	"github.com/BOPR/contracts/rollup"
	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
)

const (
	SubmitTransferMethod        = "submitTransfer"
	SubmitCreate2TransferMethod = "submitCreate2Transfer"
	SubmitMassMigrationMethod   = "submitMassMigration"
)

// SubmitBatch submits the batch on chain with updated root and compressed transactions
func (b *Bazooka) SubmitBatch(commitments []core.Commitment, accountRoot string) (txHash string, updatedCommitments []core.Commitment, err error) {
	b.log.Info(
		"Attempting to submit a new batch",
		"NumOfCommitments",
		len(commitments),
	)

	var commitmentData []core.CommitmentData

	if len(commitments) == 0 {
		b.log.Info("No transactions to submit, waiting....")
		return "", updatedCommitments, nil
	}

	switch txType := commitments[0].BatchType; txType {
	case core.TX_TRANSFER_TYPE:
		commitmentData, txHash, err = b.submitTransferBatch(commitments, accountRoot)
		if err != nil {
			return txHash, updatedCommitments, err
		}
		b.log.Info("Sent a new batch!", "TxHash", txHash, "Type", core.TX_TRANSFER_TYPE)
	case core.TX_CREATE_2_TRANSFER:
		commitmentData, txHash, err = b.submitCreate2TransferBatch(commitments, accountRoot)
		if err != nil {
			return txHash, updatedCommitments, err
		}

		b.log.Info("Sent a new batch!", "TxHash", txHash, "Type", core.TX_CREATE_2_TRANSFER)

	case core.TX_MASS_MIGRATIONS:
		commitmentData, txHash, err = b.submitMassMigrationBatch(commitments, accountRoot)
		if err != nil {
			return txHash, updatedCommitments, err
		}
		b.log.Info("Sent a new batch!", "TxHash", txHash, "Type", core.TX_MASS_MIGRATIONS)
	default:
		b.log.Error("Tx not indentified", "txType", commitments[0].BatchType)
	}

	for i := range commitments {
		commitments[i].BodyRoot = commitmentData[i].BodyRoot
		commitments[i].StateRoot = commitmentData[i].StateRoot
	}

	return txHash, commitments, nil
}

func (b *Bazooka) submitTransferBatch(commitments []core.Commitment, accountRoot string) ([]core.CommitmentData, string, error) {
	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int
	var feeReceivers []*big.Int
	var commitmentData []core.CommitmentData

	dummyReceivers := big.NewInt(0)
	for i := 0; i <= len(commitments); i++ {
		feeReceivers = append(feeReceivers, dummyReceivers)
	}

	for _, commitment := range commitments {
		compressedTxs, err := CompressTxs(b, commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return commitmentData, "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, core.BytesToByteArray(commitment.StateRoot))
		totalTxs += len(commitment.Txs)

		sig, err := core.BytesToSolSignature(commitment.AggregatedSignature)
		if err != nil {
			return commitmentData, "", err
		}
		aggregatedSig = append(aggregatedSig, sig)
	}
	b.log.Info("Batch prepared", "totalTransactions", totalTxs)

	rollupAddress := ethCmn.HexToAddress(b.Cfg.RollupAddress)
	stakeAmount := big.NewInt(int64(b.Cfg.StakeAmount))

	var inputData TransferCalldata
	inputData.StateRoots = updatedRoots
	inputData.Signatures = aggregatedSig
	inputData.FeeReceivers = feeReceivers
	inputData.Txss = txs

	commitmentData, err := inputData.Commitments(accountRoot)
	if err != nil {
		return commitmentData, "", err
	}

	tx, err := b.SignAndBroadcastBatch(b.EthClient, rollupAddress, stakeAmount, &inputData)
	if err != nil {
		return commitmentData, "", err
	}

	return commitmentData, tx.Hash().String(), nil
}

func (b *Bazooka) submitCreate2TransferBatch(commitments []core.Commitment, accountRoot string) ([]core.CommitmentData, string, error) {
	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int
	var feeReceivers []*big.Int
	var commitmentData []core.CommitmentData

	dummyReceivers := big.NewInt(0)
	for i := 0; i <= len(commitments); i++ {
		feeReceivers = append(feeReceivers, dummyReceivers)
	}

	for _, commitment := range commitments {
		compressedTxs, err := CompressTxs(b, commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return commitmentData, "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, core.BytesToByteArray(commitment.StateRoot))
		totalTxs += len(commitment.Txs)

		sig, err := core.BytesToSolSignature(commitment.AggregatedSignature)
		if err != nil {
			return commitmentData, "", err
		}
		aggregatedSig = append(aggregatedSig, sig)
	}
	b.log.Info("Batch prepared", "totalTransactions", totalTxs)

	rollupAddress := ethCmn.HexToAddress(b.Cfg.RollupAddress)
	stakeAmount := big.NewInt(int64(b.Cfg.StakeAmount))

	var inputData Create2TransferCalldata
	inputData.StateRoots = updatedRoots
	inputData.Signatures = aggregatedSig
	inputData.FeeReceivers = feeReceivers
	inputData.Txss = txs

	commitmentData, err := inputData.Commitments(accountRoot)
	if err != nil {
		return commitmentData, "", err
	}

	tx, err := b.SignAndBroadcastBatch(b.EthClient, rollupAddress, stakeAmount, inputData)
	if err != nil {
		return commitmentData, "", err
	}

	return commitmentData, tx.Hash().String(), nil
}

func (b *Bazooka) submitMassMigrationBatch(commitments []core.Commitment, accountRoot string) ([]core.CommitmentData, string, error) {
	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int

	var meta [][4]*big.Int
	var withdrawRoots [][32]byte
	var commitmentData []core.CommitmentData

	dummyReceiver := big.NewInt(0)

	for _, commitment := range commitments {
		compressedTxs, err := CompressTxs(b, commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return commitmentData, "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, core.BytesToByteArray(commitment.StateRoot))
		totalTxs += len(commitment.Txs)

		sig, err := core.BytesToSolSignature(commitment.AggregatedSignature)
		if err != nil {
			return commitmentData, "", err
		}

		aggregatedSig = append(aggregatedSig, sig)

		var spokeID = big.NewInt(0)
		var tokenID = big.NewInt(0)
		var totalAmount = big.NewInt(0)

		for i, tx := range commitment.Txs {
			_, spoke, _, _, amount, _, err := b.DecodeMassMigrationTx(tx.Data)
			if err != nil {
				return commitmentData, "", err
			}

			if i == 0 {
				spokeID = spoke
				// TODO fix
				// state, err := core.DBInstance.GetStateByIndex(from.Uint64())
				// if err != nil {
				// 	return "", err
				// }
				// _, _, _, token, err := b.DecodeState(state.Data)
				// if err != nil {
				// 	return "", err
				// }
				tokenID = big.NewInt(0)
			}

			totalAmount.Add(amount, totalAmount)
		}

		var metaValues [4]*big.Int
		metaValues[0] = spokeID
		metaValues[1] = tokenID
		metaValues[2] = totalAmount
		metaValues[3] = dummyReceiver

		meta = append(meta, metaValues)
	}

	withdrawRoots = updatedRoots

	b.log.Info("Batch prepared", "totalTransactions", totalTxs)

	rollupAddress := ethCmn.HexToAddress(b.Cfg.RollupAddress)
	stakeAmount := big.NewInt(int64(b.Cfg.StakeAmount))

	var inputData MassMigrationCalldata
	inputData.StateRoots = updatedRoots
	inputData.WithdrawRoots = withdrawRoots
	inputData.Signatures = aggregatedSig
	inputData.Meta = meta
	inputData.Txss = txs

	commitmentData, err := inputData.Commitments(accountRoot)
	if err != nil {
		return commitmentData, "", err
	}

	tx, err := b.SignAndBroadcastBatch(b.EthClient, rollupAddress, stakeAmount, inputData)
	if err != nil {
		return commitmentData, "", err
	}

	return commitmentData, tx.Hash().String(), nil
}

func (b *Bazooka) FireDepositFinalisation(TBreplaced core.UserState, siblings []core.UserState, commitmentMP TypesCommitmentInclusionProof, subTreeHeight uint64) (err error) {
	b.log.Info(
		"Attempting to finalise deposits",
		"NodeToBeReplaced",
		TBreplaced,
		"NumberOfSiblings",
		len(siblings),
		"atDepth",
		subTreeHeight,
	)

	stateProof := NewStateMerkleProof(TBreplaced, siblings)
	solStateProof, err := stateProof.ToABIVersion(*b)
	if err != nil {
		return err
	}
	pathAtDepth := core.StringToBigInt(TBreplaced.Path)

	commitmentIP := rollup.TypesCommitmentInclusionProof{
		Commitment: rollup.TypesCommitment{
			StateRoot: core.BytesToByteArray(commitmentMP.Commitment.StateRoot), BodyRoot: core.BytesToByteArray(commitmentMP.Commitment.BodyRoot),
		},
		Path:    commitmentMP.Path,
		Witness: commitmentMP.Witness,
	}

	vacancyProof := rollup.TypesSubtreeVacancyProof{
		Depth:       big.NewInt(int64(subTreeHeight)),
		Witness:     solStateProof.Witness,
		PathAtDepth: pathAtDepth,
	}

	input, err := b.RollupABI.Pack("submitDeposits", commitmentIP, vacancyProof)
	if err != nil {
		b.log.Error("Error packing data for submitBatch", "err", err)
		return err
	}

	stakeAmount := big.NewInt(int64(b.Cfg.StakeAmount))
	tx, err := b.SignAndBroadcast(b.EthClient, ethCmn.HexToAddress(b.Cfg.RollupAddress), stakeAmount, input)
	if err != nil {
		b.log.Error("Error sending register batch", "err", err)
		return
	}

	b.log.Info("Deposits successfully finalized!", "TxHash", tx.Hash())
	return nil
}

// RegisterPubkeys registers pubkeys in a batch
func (b *Bazooka) RegisterPubkeys(pubkeys [16][4]*big.Int) (txHash string, err error) {
	registryABI, err := abi.JSON(strings.NewReader(accountregistry.AccountregistryABI))
	if err != nil {
		return
	}
	data, err := registryABI.Pack("registerBatch", pubkeys)
	if err != nil {
		b.log.Error("Error packing data for register batch", "err", err)
		return
	}

	tx, err := b.SignAndBroadcast(b.EthClient, ethCmn.HexToAddress(b.Cfg.AccountRegistry), big.NewInt(0), data)
	if err != nil {
		b.log.Error("Error sending register batch", "err", err)
		return
	}

	b.log.Info("Registered pubkeys", "count", len(pubkeys), "txHash", tx.Hash().String())

	return tx.Hash().String(), nil
}
