#!/bin/bash

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=arm
export GOARM=6

#go build -trimpath -ldflags '-s -w -X github.com/sweatybridge/pijoy/internal.version=dev'
#scp ./pijoy raspberrypi.local:/home/qh812
# go build -trimpath -ldflags '-s -w -X github.com/sweatybridge/pijoy/internal.version=dev' tools/debug/main.go
# scp ./main raspberrypi.local:/home/qh812
go build -trimpath -ldflags '-s -w -X github.com/sweatybridge/pijoy/internal.version=dev' tools/console/main.go
scp ./main qh812@raspberrypi.local:/home/qh812
