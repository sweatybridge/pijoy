#!/bin/bash
set -euo pipefail

cmd="sudo systemctl enable ${1:-pijoy}"

# shellcheck disable=SC2029
ssh qh812@raspberrypi.local "$cmd"
