#!/bin/bash
set -euo pipefail

cmd="journalctl -u ${1:-pijoy} -f"

# shellcheck disable=SC2029
ssh qh812@raspberrypi.local "$cmd"
