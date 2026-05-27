#!/usr/bin/env sh
set -eu

cd "$(dirname "$0")"
go build -o agentx .

printf '%s\n' "Built helper/agentx"
