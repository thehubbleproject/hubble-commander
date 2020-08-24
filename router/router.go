package router

import (
	"github.com/BOPR/core"
)

type HubbleTxBinders interface {
	compressTxs(txs []core.Tx) ([]byte, error)
	processTx(balanceRoot, accountRoot core.ByteArray) (newBalanceRoot core.ByteArray, accounts [][]byte, err error)
	applyTxWithProof(accountMP core.AccountMerkleProof, tx core.Tx) (updatedAccount []byte, updatedRoot core.ByteArray, err error)
	applyTxWithoutProof(tx core.Tx) (updatedAccount []byte, updatedRoot core.ByteArray, err error)
}

type Router struct {
	registeredRoutes map[int]HubbleTxBinders
}

func NewRouter() *Router {
	routes := make(map[int]HubbleTxBinders)
	return &Router{registeredRoutes: routes}
}

func (r *Router) RegisterNewTransaction(id int, transactionBinders HubbleTxBinders) {
	r.registeredRoutes[id] = transactionBinders
}

func (r *Router) ProcessTx(tx core.Tx) (newBalanceRoot core.ByteArray, accounts [][]byte, err error) {
	return r.registeredRoutes[int(tx.Type)].processTx()
}

func (r *Router) CompressTxs(txs []core.Tx, txType int) ([]byte, error) {
	return r.registeredRoutes[txType].compressTxs(txs)
}

func (r *Router) ApplyTx(tx core.Tx) (updatedAccount []byte, updatedRoot core.ByteArray, err error) {
	return r.registeredRoutes[int(tx.Type)].applyTxWithoutProof(tx)
}

// TODO need to create functions for sign bytes and decompress
