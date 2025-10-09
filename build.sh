set -euxo pipefail

# Get the directory that this script file is in
THIS_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

cd "$THIS_DIR"

mkdir -p "$THIS_DIR/functions"
LDFLAGS="-linkmode external -extldflags "-static""
CGO_ENABLED=0 GOBIN=$THIS_DIR/functions go install -ldflags "$LDFLAGS" ./...
chmod +x "$THIS_DIR"/functions/*
go env
