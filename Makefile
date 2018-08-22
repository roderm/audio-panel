generate-proto:
	protoc -I. -I$$GOPATH/src/ ./proto/service.proto \
		--go_out=plugins=grpc:$$GOPATH/src/
js-gen:
	NODE_ENV=production webpack -p --progress
dev:
	NODE_ENV=development webpack -d --watch --inspect &
	go run main.go