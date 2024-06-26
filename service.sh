#!/bin/bash
set -euo pipefail

if [[ "${1:-}" == 'update' ]]; then
    rsync --rsync-path="sudo rsync" pijoy.service qh812@raspberrypi.local:/etc/systemd/system
    ssh qh812@raspberrypi.local 'sudo systemctl daemon-reload'
fi

echo 'Restarting PiJoy API server...'
ssh qh812@raspberrypi.local 'sudo systemctl restart pijoy'
