build-rpi:
	mkdir -p ../../../build/rpi/drivers/
	env GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED="1" CC=arm-linux-gnueabi-gcc go build -buildmode=plugin -o ../../../build/rpi/drivers/pioneer.so -a
build:
	mkdir -p ../../../build/linux/drivers/
	env GOOS=linux go build -buildmode=plugin -o ../../../build/linux/drivers/pioneer.so