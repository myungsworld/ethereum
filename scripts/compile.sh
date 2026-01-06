#!/bin/bash
# Solidity 컴파일 및 Go 바인딩 생성 스크립트

set -e

# 경로 설정
ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
SOLIDITY_DIR="$ROOT_DIR/solidity"
CONTRACTS_DIR="$SOLIDITY_DIR/contracts"
BUILD_DIR="$SOLIDITY_DIR/build"
GO_BINDINGS_DIR="$ROOT_DIR/go/internal/infrastructure/contracts"

echo "=== Solidity 컴파일 시작 ==="
echo "Root: $ROOT_DIR"

# 빌드 디렉토리 생성
mkdir -p "$BUILD_DIR"
mkdir -p "$GO_BINDINGS_DIR"

# 모든 .sol 파일 컴파일
for sol_file in "$CONTRACTS_DIR"/*.sol; do
    if [ -f "$sol_file" ]; then
        filename=$(basename "$sol_file" .sol)
        echo "컴파일 중: $filename.sol"

        # solc로 ABI와 BIN 생성
        solc --abi --bin --overwrite \
            --optimize \
            -o "$BUILD_DIR" \
            "$sol_file"

        echo "  → $BUILD_DIR/$filename.abi"
        echo "  → $BUILD_DIR/$filename.bin"

        # abigen으로 Go 바인딩 생성
        abi_file="$BUILD_DIR/$filename.abi"
        bin_file="$BUILD_DIR/$filename.bin"

        if [ -f "$abi_file" ] && [ -f "$bin_file" ]; then
            go_file="$GO_BINDINGS_DIR/$(echo "$filename" | tr '[:upper:]' '[:lower:]').go"

            abigen --abi="$abi_file" \
                   --bin="$bin_file" \
                   --pkg=contracts \
                   --type="$filename" \
                   --out="$go_file"

            echo "  → $go_file"
        fi
    fi
done

echo ""
echo "=== 컴파일 완료 ==="
echo "ABI/BIN: $BUILD_DIR"
echo "Go 바인딩: $GO_BINDINGS_DIR"
