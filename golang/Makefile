all: init proto-gen

proto-compiler:
	curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
	unzip protoc-3.15.8-linux-x86_64.zip -d "${HOME}/.local"
	export PATH="${PATH}:${HOME}/.local/bin"

init:
	@echo "Pull External Protofiles"
	#cd ../
	#git submodule init
	#git submodule update --remote
	#cd python

proto-gen:
	@echo "Generating Protobuf files"
	#yum install  protobuf-compiler
	./scripts/proto-gen.sh

.PHONY: all proto-gen init
