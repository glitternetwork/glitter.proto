all: init proto-gen build

proto-compiler:
	curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
	unzip protoc-3.15.8-linux-x86_64.zip -d "${HOME}/.local"
	export PATH="${PATH}:${HOME}/.local/bin"

init:
	@echo "Pull External Protofiles"
	git submodule init
	git submodule update --remote
	rm -rf ../proto/index
	cp -R ../glitter-chain/proto/index ../proto

proto-gen:
	@echo "Generating Protobuf files"
	#yum install  protobuf-compiler
	./scripts/proto-gen.sh

build: 
	pip install poetry
	poetry build

publish:
	poetry publish 

.PHONY: all proto-gen format init build publish
