build-hubble-contracts:
	cd .contracts/ && npm ci && npm run generate

run-hardhat-node:
	cd .contracts/ && npm run node

deploy-hubble-contracts:
	cd .contracts/ && npm run deploy

build-bindings:
	go install github.com/ethereum/go-ethereum/cmd/abigen
	go generate

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
