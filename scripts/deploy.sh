#!/bin/bash
set -euo pipefail

# goreleaser build --snapshot --clean

if [[ "${1:-}" == 'init' ]]; then
    echo "Deploying service init files..."
    rsync -a --rsync-path="sudo rsync" init/ qh812@raspberrypi.local:/etc/systemd/system
    ssh qh812@raspberrypi.local 'sudo systemctl daemon-reload'
else
    echo "Deploying service binary files..."
    # scp fails due to write protection of systemd service
    rsync -a dist/pizero/ qh812@raspberrypi.local:/home/qh812/pijoy
fi
