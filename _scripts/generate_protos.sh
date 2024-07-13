#!/bin/bash

# Gerar os arquivos gRPC Go a partir de um arquivo proto.
# params:
# $1: diretório do arquivo proto
# $2: valor a ser substituído no caminho do arquivo gerado
#
# example:
# ./_scripts/generate_protos.sh ./proto user

generate_protos() {
    local proto_dir="$1"
    local replace_value="$2"

    output_dir="./internal/${replace_value}/adapter/grpc/generated"
    mkdir -p "$output_dir"

    protoc --go_out="$output_dir" \
           --go-grpc_out="$output_dir" \
           -I "$proto_dir" \
           --go_opt=paths=source_relative \
           --go-grpc_opt=paths=source_relative \
           "$proto_dir/${replace_value}.proto"
}

if [ $# -lt 2 ]; then
    echo "Usage: $0 <proto_dir> <replace_value>"
    exit 1
fi

generate_protos "$1" "$2"
