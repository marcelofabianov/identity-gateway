#!/bin/bash

# Diretório dos arquivos proto
PROTO_DIR="./api/v1/proto"
# Diretório para os arquivos gerados
GEN_DIR="./api/v1/gen"

# Certifique-se de que o diretório de destino existe
mkdir -p $GEN_DIR

# Encontre todos os arquivos .proto no diretório PROTO_DIR
PROTO_FILES=$(find $PROTO_DIR -name "*.proto")

# Gere os arquivos Go para cada arquivo .proto
for FILE in $PROTO_FILES; do
    echo "Generating Go code for $FILE"
    protoc \
        --go_out=$GEN_DIR --go_opt=paths=source_relative \
        --go-grpc_out=$GEN_DIR --go-grpc_opt=paths=source_relative \
        --proto_path=$PROTO_DIR $FILE
done

echo "All proto files have been processed."
