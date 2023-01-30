# go build
.PHONY: build
build:
	GOOS=linux GOARCH=arm GOARM=7 go build -o mimamori-exec-server

# scp upload
.PHONY: scp
scp:
	scp ./mimamori-exec-server pi@192.168.0.100:/home/pi/mimamori
