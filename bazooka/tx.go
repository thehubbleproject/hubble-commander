package bazooka

import (
	"math/big"
	"strings"

	"github.com/BOPR/config"
	"github.com/BOPR/contracts/accountregistry"
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
func (b *Bazooka) SubmitBatch(commitments []core.Commitment) (txHash string, err error) {
	b.log.Info(
		"Attempting to submit a new batch",
		"NumOfCommitments",
		len(commitments),
	)

	if len(commitments) == 0 {
		b.log.Info("No transactions to submit, waiting....")
		return "", nil
	}

	switch txType := commitments[0].BatchType; txType {
	case core.TX_TRANSFER_TYPE:
		txHash, err := b.submitTransferBatch(commitments)
		if err != nil {
			return "", err
		}
		b.log.Info("Sent a new batch!", "TxHash", txHash, "Type", core.TX_TRANSFER_TYPE)
	case core.TX_CREATE_2_TRANSFER:
		txHash, err := b.submitCreate2TransferBatch(commitments)
		if err != nil {
			return "", err
		}
		b.log.Info("Sent a new batch!", "TxHash", txHash, "Type", core.TX_CREATE_2_TRANSFER)
	case core.TX_MASS_MIGRATIONS:
		txHash, err := b.submitMassMigrationBatch(commitments)
		if err != nil {
			return "", err
		}
		b.log.Info("Sent a new batch!", "TxHash", txHash, "Type", core.TX_MASS_MIGRATIONS)
	default:
		b.log.Error("Tx not indentified", "txType", commitments[0].BatchType)
	}

	return txHash, nil
}

func (b *Bazooka) submitTransferBatch(commitments []core.Commitment) (string, error) {
	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int
	var feeReceivers []*big.Int

	dummyReceivers := big.NewInt(0)
	for i := 0; i <= len(commitments); i++ {
		feeReceivers = append(feeReceivers, dummyReceivers)
	}

	for _, commitment := range commitments {
		compressedTxs, err := CompressTxs(b, commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, commitment.UpdatedRoot)
		totalTxs += len(commitment.Txs)

		sig, err := core.BytesToSolSignature(commitment.AggregatedSignature)
		if err != nil {
			return "", err
		}
		aggregatedSig = append(aggregatedSig, sig)
	}
	b.log.Info("Batch prepared", "totalTransactions", totalTxs)

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)

	// TODO https://github.com/thehubbleproject/hubble-commander/issues/68
	stakeAmount := big.NewInt(1000000000000000000)

	var inputData TransferCalldata
	inputData.StateRoots = updatedRoots
	inputData.Signatures = aggregatedSig
	inputData.FeeReceivers = feeReceivers
	inputData.Txss = txs

	tx, err := b.SignAndBroadcastBatch(b.EthClient, rollupAddress, stakeAmount, &inputData)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

func (b *Bazooka) submitCreate2TransferBatch(commitments []core.Commitment) (string, error) {
	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int
	var feeReceivers []*big.Int

	dummyReceivers := big.NewInt(0)
	for i := 0; i <= len(commitments); i++ {
		feeReceivers = append(feeReceivers, dummyReceivers)
	}

	for _, commitment := range commitments {
		compressedTxs, err := CompressTxs(b, commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, commitment.UpdatedRoot)
		totalTxs += len(commitment.Txs)

		sig, err := core.BytesToSolSignature(commitment.AggregatedSignature)
		if err != nil {
			return "", err
		}
		aggregatedSig = append(aggregatedSig, sig)
	}
	b.log.Info("Batch prepared", "totalTransactions", totalTxs)

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)

	// TODO https://github.com/thehubbleproject/hubble-commander/issues/68
	stakeAmount := big.NewInt(1000000000000000000)

	var inputData Create2TransferCalldata
	inputData.StateRoots = updatedRoots
	inputData.Signatures = aggregatedSig
	inputData.FeeReceivers = feeReceivers
	inputData.Txss = txs

	tx, err := b.SignAndBroadcastBatch(b.EthClient, rollupAddress, stakeAmount, inputData)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

func (b *Bazooka) submitMassMigrationBatch(commitments []core.Commitment) (string, error) {
	var txs [][]byte
	var updatedRoots [][32]byte
	var aggregatedSig [][2]*big.Int
	var totalTxs int

	var meta [][4]*big.Int
	var withdrawRoots [][32]byte

	dummyReceiver := big.NewInt(0)

	for _, commitment := range commitments {
		compressedTxs, err := CompressTxs(b, commitment.Txs)
		if err != nil {
			b.log.Error("Unable to compress txs", "error", err)
			return "", err
		}
		txs = append(txs, compressedTxs)
		updatedRoots = append(updatedRoots, commitment.UpdatedRoot)
		totalTxs += len(commitment.Txs)

		sig, err := core.BytesToSolSignature(commitment.AggregatedSignature)
		if err != nil {
			return "", err
		}

		aggregatedSig = append(aggregatedSig, sig)

		var spokeID = big.NewInt(0)
		var tokenID = big.NewInt(0)
		var totalAmount = big.NewInt(0)

		for i, tx := range commitment.Txs {
			_, spoke, _, _, amount, _, err := b.DecodeMassMigrationTx(tx.Data)
			if err != nil {
				return "", err
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

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)

	// TODO https://github.com/thehubbleproject/hubble-commander/issues/68
	stakeAmount := big.NewInt(100000000000000000)

	var inputData MassMigrationCalldata
	inputData.StateRoots = updatedRoots
	inputData.WithdrawRoots = withdrawRoots
	inputData.Signatures = aggregatedSig
	inputData.Meta = meta
	inputData.Txss = txs

	tx, err := b.SignAndBroadcastBatch(b.EthClient, rollupAddress, stakeAmount, inputData)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}

func (b *Bazooka) FireDepositFinalisation(TBreplaced core.UserState, siblings []core.UserState, subTreeHeight uint64) (err error) {
	b.log.Info(
		"Attempting to finalise deposits",
		"NodeToBeReplaced",
		TBreplaced,
		"NumberOfSiblings",
		len(siblings),
		"atDepth",
		subTreeHeight,
	)

	// depositSubTreeHeight := big.NewInt(0)
	// depositSubTreeHeight.SetUint64(subTreeHeight)
	// var siblingData [][32]byte
	// for _, sibling := range siblings {
	// 	data, err := HexToByteArray(sibling.Hash)
	// 	if err != nil {
	// 		b.log.Error("unable to convert HexToByteArray", err)
	// 		return err
	// 	}
	// 	siblingData = append(siblingData, data)
	// }

	// accountProof := rollup.TypesAccountMerkleProof{}
	// accountProof.AccountIP.PathToAccount = StringToBigInt(TBreplaced.Path)
	// userAccount, err := TBreplaced.ToABIAccount()
	// if err != nil {
	// 	b.log.Error("unable to convert", "error", err)
	// 	return
	// }
	// accountProof.AccountIP.Account = rollup.TypesUserState(userAccount)

	// accountProof.Siblings = siblingData
	// data, err := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Pack("finaliseDepositsAndSubmitBatch", depositSubTreeHeight, accountProof)
	// if err != nil {
	// 	fmt.Println("Unable to craete data", err)
	// 	return
	// }

	// rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	// stakeAmount := big.NewInt(0)
	// stakeAmount.SetString("32000000000000000000", 10)

	// // generate call msg
	// callMsg := ethereum.CallMsg{
	// 	To:    &rollupAddress,
	// 	Data:  data,
	// 	Value: stakeAmount,
	// }

	// auth, err := b.generateAuthObj(b.EthClient, callMsg)
	// if err != nil {
	// 	return err
	// }
	// b.log.Info("Broadcasting deposit finalisation transaction")

	// tx, err := b.RollupContract.FinaliseDepositsAndSubmitBatch(auth, depositSubTreeHeight, accountProof)
	// if err != nil {
	// 	return err
	// }
	// b.log.Info("Deposits successfully finalized!", "TxHash", tx.Hash())
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

	tx, err := b.SignAndBroadcast(b.EthClient, ethCmn.HexToAddress(config.GlobalCfg.AccountRegistry), big.NewInt(0), data)
	if err != nil {
		b.log.Error("Error sending register batch", "err", err)
		return
	}

	b.log.Info("Registered pubkeys", "count", len(pubkeys), "txHash", tx.Hash().String())

	return tx.Hash().String(), nil
}
