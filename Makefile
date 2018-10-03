pldirs = plugin/drivers/pioneer plugin/drivers/nanoleaf

test: 
	ls -l $(pldirs)
build-rpi:
	rm -rf build/rpi/
	mkdir -p build/rpi/commandsets/
	env GOOS=linux GOARCH=arm GOARM=5 go build -o build/rpi/audio-panel
	make js-gen
	make build-plugins-rpi
	cp -r commandsets/* build/rpi/commandsets/

build-linux:
	rm -rf build/linux/
	mkdir -p build/linux/commandsets/
	env GOOS=linux GOARCH=amd64 go build -o build/linux/audio-panel
	make js-gen
	make build-plugins
	cp -r commandsets/* build/linux/commandsets/

build-plugins-rpi:
	$(foreach dir,$(pldirs),(cd $(dir) && make build-rpi ) &&) :

build-plugins:
	$(foreach dir,$(pldirs),(cd $(dir) && make build ) &&) :

generate-proto:
	protoc -I. -I$$GOPATH/src/ ./proto/service.proto \
		--go_out=plugins=grpc:$$GOPATH/src/
	ls ./proto/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
dev:
	go run main.go ./dev-config.json