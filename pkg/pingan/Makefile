
GOPATH:=$(shell go env GOPATH)
OBJECTDIR:=$(shell dirname $(shell dirname $(shell pwd)))
PROTO = pingan

.PHONY: proto
proto:
	@echo execute ${PROTO} proto file generate ${OBJECTDIR}
	protoc --proto_path=. --proto_path=${GOPATH}/src  --go_out=. proto/${PROTO}.proto
