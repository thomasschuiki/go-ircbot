 #Variables
MAJOR_VERSION := 1
MINOR_VERSION := 0
PATCH_VERSION := 0
BUILD := $(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags "-X=main.version=${MAJOR_VERSION}.${MINOR_VERSION}.${PATCH_VERSION}-${BUILD} -s -w"
BINARY := go-ircbot

build:
	go build ${LDFLAGS} -o "bin/${BINARY}_${MAJOR_VERSION}.${MINOR_VERSION}.${PATCH_VERSION}" main.go

run:
	go run main.go

debug:
	IRCCHANNELS="#uadachannel" IRCNICK="juicybot" IRCSERVER="irc.quakenet.org:6667" IRCUSER="notchloe" go run main.go

dev:
	go run debug/main.go
