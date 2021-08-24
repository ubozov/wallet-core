protogen:
	protoc --proto_path=./proto/v1/  --go_out=plugins=grpc:./proto/v1 ./proto/v1/*.proto

run: build
	go run . --config ./go-webapi.config.json

build: protogen
	go build .