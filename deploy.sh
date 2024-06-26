#!/bin/bash
set -euo pipefail

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=arm
export GOARM=6

SRC=${1:-.}
BIN=$(basename "$(dirname "$SRC")")
if [[ "$BIN" == '.' ]]; then
    BIN='pijoy'
fi

go build -trimpath -ldflags '-s -w -X github.com/sweatybridge/pijoy/internal.version=dev' -o "$BIN" "$SRC"

# scp fails due to write protection of systemd service
rsync "$BIN" qh812@raspberrypi.local:/home/qh812
