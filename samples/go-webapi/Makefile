protogen:
	protoc --proto_path=./crypto/proto/  --go_out=plugins=grpc:./crypto/proto ./crypto/proto/*.proto

run: build
	go run . --config ./go-webapi.config.json

build: protogen
	go build .