# go build
.PHONY: build
build:
	GOOS=linux GOARCH=arm GOARM=7 go build -o mimamori-camera-controller

# scp upload
.PHONY: scp
scp:
	scp ./mimamori-camera-controller pi@192.168.0.10:/home/pi/mimamori
