// Package gen generates abi and contract bindings.
package gen

//go:generate sh -c "rm -rf abi/ && mkdir abi/"

//go:generate sh -c "cat .contracts/artifacts/contracts/rollup/Rollup.sol/Rollup.json					| jq .abi -M > abi/rollup.abi"
//go:generate sh -c "cat .contracts/artifacts/contracts/client/FrontendGeneric.sol/FrontendGeneric.json			| jq .abi -M > abi/state.abi"
//go:generate sh -c "cat .contracts/artifacts/contracts/DepositManager.sol/DepositManager.json				| jq .abi -M > abi/depositmanager.abi"
//go:generate sh -c "cat .contracts/artifacts/contracts/TokenRegistry.sol/TokenRegistry.json				| jq .abi -M > abi/tokenregistry.abi"
//go:generate sh -c "cat .contracts/artifacts/contracts/proposers/BurnAuction.sol/BurnAuction.json			| jq .abi -M > abi/burnauction.abi"
//go:generate sh -c "cat .contracts/artifacts/contracts/client/FrontendTransfer.sol/FrontendTransfer.json		| jq .abi -M > abi/transfer.abi"
//go:generate sh -c "cat .contracts/artifacts/contracts/client/FrontendMassMigration.sol/FrontendMassMigration.json	| jq .abi -M > abi/massmigration.abi"
//go:generate sh -c "cat .contracts/artifacts/contracts/client/FrontendCreate2Transfer.sol/FrontendCreate2Transfer.json	| jq .abi -M > abi/create2transfer.abi"
//go:generate sh -c "cat .contracts/artifacts/contracts/BLSAccountRegistry.sol/BLSAccountRegistry.json			| jq .abi -M > abi/accountregistry.abi"

//go:generate rm -rf contracts/
//go:generate mkdir -p contracts/rollup
//go:generate mkdir -p contracts/state
//go:generate mkdir -p contracts/depositmanager
//go:generate mkdir -p contracts/tokenregistry
//go:generate mkdir -p contracts/burnauction
//go:generate mkdir -p contracts/transfer
//go:generate mkdir -p contracts/massmigration
//go:generate mkdir -p contracts/create2transfer
//go:generate mkdir -p contracts/accountregistry

//go:generate abigen --abi abi/rollup.abi		--pkg=rollup		--out=contracts/rollup/rollup.go
//go:generate abigen --abi abi/state.abi		--pkg=state		--out=contracts/state/state.go
//go:generate abigen --abi abi/depositmanager.abi	--pkg=depositmanager	--out=contracts/depositmanager/depositmanager.go
//go:generate abigen --abi abi/tokenregistry.abi	--pkg=tokenregistry	--out=contracts/tokenregistry/tokenregistry.go
//go:generate abigen --abi abi/burnauction.abi		--pkg=burnauction	--out=contracts/burnauction/burnauction.go
//go:generate abigen --abi abi/transfer.abi		--pkg=transfer		--out=contracts/transfer/transfer.go
//go:generate abigen --abi abi/massmigration.abi	--pkg=massmigration	--out=contracts/massmigration/massmigration.go
//go:generate abigen --abi abi/create2transfer.abi	--pkg=create2transfer	--out=contracts/create2transfer/create2transfer.go
//go:generate abigen --abi abi/accountregistry.abi	--pkg=accountregistry	--out=contracts/accountregistry/accountregistry.go
