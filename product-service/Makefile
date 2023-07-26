PROTO_SRC_PATH := ./tools/protos

run:
	go run $$(find cmd/ -name main.go)

test:
	go test -tags='test' ./...

protobuf-product:
	protoc --go_out=. $(PROTO_SRC_PATH)/product.proto
	protoc --go-grpc_out=. $(PROTO_SRC_PATH)/product.proto
