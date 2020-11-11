dep:
	dep ensure -v
	rm -rf vendor/github.com/ethereum
	mkdir -p vendor/github.com/ethereum
	git clone -b v1.9.0 --single-branch --depth 1 https://github.com/ethereum/go-ethereum vendor/github.com/ethereum/go-ethereum
	
contracts:
	abigen --abi=contracts/rollup/rollup.abi --pkg=rollup --out=contracts/rollup/rollup.go
	abigen --abi=contracts/registry/registry.abi --pkg=registry --out=contracts/registry/registry.go
	abigen --abi=contracts/logger/logger.abi --pkg=logger --out=contracts/logger/logger.go
	abigen --abi=contracts/rollupclient/rollupclient.abi --pkg=rollupclient --out=contracts/rollupclient/rollupclient.go

clean:
	rm -rf build

build: clean
	mkdir -p build
	go build -o build/hubble ./cmd

buidl: build

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

.PHONY: contracts dep start-simulator build clean start buidl
