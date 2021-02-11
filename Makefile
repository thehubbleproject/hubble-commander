build-hubble-contracts:
	cd .contracts/ && npm ci && npm run generate

run-hardhat-node:
	mkdir -p ./logs
	cd .contracts/ && npm run node >> ../logs/node.log 2>&1 &

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
	docker stop mysql || true && docker rm mysql || true
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

get-submodules:
	git submodule init && git submodule update -r

setup: build init load create-database migrate-up reset

setup-integration: run-database get-submodules build-hubble-contracts build-bindings run-hardhat-node deploy-hubble-contracts setup

start:
	mkdir -p logs &
	touch ./logs/node.log
	./build/hubble start > ./logs/node.log & 

.PHONY: contracts dep start-simulator build clean start buidl lint
