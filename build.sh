set -euxo pipefail

THIS_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

cd "$THIS_DIR"

go env
go mod download

FUNC_DIR=$THIS_DIR/functions
mkdir -p "$FUNC_DIR"
GOBIN=$FUNC_DIR go install ./...
