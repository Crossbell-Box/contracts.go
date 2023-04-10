all: clean build/contract

clean: clean-build

build/contract:
	mkdir -p build/tmp/
	cd Crossbell-Contracts/ && \
	yarn && \
	make install && \
	make build && \
	forge inspect Web3Entry abi > ../build/tmp/Web3Entry.abi && \
	forge inspect Web3Entry b > ../build/tmp/Web3Entry.bin && \
	cd ../ && \
	mkdir -p build/contract/
	abigen --bin=build/tmp/Web3Entry.bin --abi=build/tmp/Web3Entry.abi --pkg=contracts --out=build/contract/Web3Entry.go

clean-build:
	rm -rf ./build/*
