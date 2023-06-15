set -euxo pipefail

mkdir -p "$(pwd)/functions"
GOBIN=$(pwd)/functions go install github.com/spotlightpa/almanack/funcs/almanack-api@latest
chmod +x "$(pwd)"/functions/*
go env
