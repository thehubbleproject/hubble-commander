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

run-database:
	docker run --name=mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql

init:
	./build/hubble init

load:
	./build/hubble load

create-database:
	./build/hubble create-database

reset:
	./build/hubble migration down --all
	./build/hubble migration up

migrate-up:
	./build/hubble migration up

migrate-down:
	./build/hubble migration down --all

setup: build init load create-database migrate-up

start:
	mkdir -p logs &
	touch ./logs/node.log
	./build/hubble start > ./logs/node.log & 

.PHONY: contracts dep start-simulator build clean start buidl lint
