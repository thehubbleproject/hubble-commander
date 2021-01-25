contracts:
	abigen --abi=contracts/rollup/rollup.abi --pkg=rollup --out=contracts/rollup/rollup.go
	abigen --abi=contracts/registry/registry.abi --pkg=registry --out=contracts/registry/registry.go
	abigen --abi=contracts/state/state.abi --pkg=state --out=contracts/state/state.go
	abigen --abi=contracts/depositmanager/depositmanager.abi --pkg=depositmanager --out=contracts/depositmanager/depositmanager.go
	abigen --abi=contracts/tokenregistry/tokenregistry.abi --pkg=tokenregistry --out=contracts/tokenregistry/tokenregistry.go
	abigen --abi=contracts/burnauction/burnauction.abi --pkg=burnauction --out=contracts/burnauction/burnauction.go
	abigen --abi=contracts/erc20/erc20.abi --pkg=erc20 --out=contracts/erc20/erc20.go
	abigen --abi=contracts/transfer/transfer.abi --pkg=transfer --out=contracts/transfer/transfer.go
	abigen --abi=contracts/massmigration/massmigration.abi --pkg=massmigration --out=contracts/massmigration/massmigration.go
	abigen --abi=contracts/create2transfer/create2transfer.abi --pkg=create2transfer --out=contracts/create2transfer/create2transfer.go
	abigen --abi=contracts/accountregistry/accountregistry.abi --pkg=accountregistry --out=contracts/accountregistry/accountregistry.go

clean:
	rm -rf build

build: clean
	mkdir -p build
	go build -o build/hubble ./cmd

buidl: build

lint:
	golangci-lint run ./...

init:
	./build/hubble init

reset:
	./build/hubble migration down --all
	./build/hubble migration up

migrate-up:
	./build/hubble migration up

migrate-down:
	./build/hubble migration down --all

start:
	mkdir -p logs &
	touch ./logs/node.log
	./build/hubble start > ./logs/node.log & 

.PHONY: contracts dep start-simulator build clean start buidl lint
