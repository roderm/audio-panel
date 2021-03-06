build-rpi:
	rm -rf build/rpi/
	mkdir -p build/rpi/commandsets/ build/rpi/views
	env GOOS=linux GOARCH=arm GOARM=5 go build -o build/rpi/audio-panel
	make js-gen
	cp -r commandsets/* build/rpi/commandsets/
	cp -r views/* build/rpi/views/

generate-proto:
	protoc -I. -I$$GOPATH/src/ ./proto/service.proto \
		--go_out=plugins=grpc:$$GOPATH/src/
	ls ./proto/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
js-gen:
	NODE_ENV=production webpack -p --progress --optimize-minimize
dev:
	NODE_ENV=development webpack -d --watch --inspect --watch-poll &
	go run main.go ./config.json