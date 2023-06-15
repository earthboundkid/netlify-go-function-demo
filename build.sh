set -euxo pipefail

mkdir -p "$(pwd)/functions"
GOBIN=$(pwd)/functions go install ./cmd/...
chmod +x "$(pwd)"/functions/*
go env
