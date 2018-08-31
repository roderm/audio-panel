generate-proto:
	protoc -I. -I$$GOPATH/src/ ./proto/service.proto \
		--go_out=plugins=grpc:$$GOPATH/src/
	ls ./proto/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
js-gen:
	NODE_ENV=production webpack -p --progress
dev:
	NODE_ENV=development webpack -d --watch --inspect &
	go run main.go ./config.json