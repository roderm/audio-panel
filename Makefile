generate-proto:
	protoc -I. -I$$GOPATH/src/ ./proto/service.proto \
		--go_out=plugins=grpc:$$GOPATH/src/
js-gen:
	NODE_ENV=production webpack -p --progress
dev:
	make js-gen
	go run main.go