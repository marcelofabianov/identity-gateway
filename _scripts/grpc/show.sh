#!/bin/bash

echo "________________________________________________________________________"

# Verificar se o nome do serviço foi fornecido como argumento
if [ "$#" -ne 1 ]; then
    echo "Uso: $0 <nome-do-serviço>"
    exit 1
fi

# Nome do serviço a ser detalhado
SERVICE_NAME=$1

# Endereço do servidor gRPC
ADDRESS="localhost:50051"

# Caminho para os certificados e chaves (ajuste conforme necessário)
CERT_DIR="./storage/certs"

# Caminho para os arquivos .proto (ajuste conforme necessário)
PROTO_DIR="./api/v1/proto"

# Comando grpcurl para listar todos os métodos disponíveis para o serviço especificado
grpcurl -v \
    -cacert $CERT_DIR/ca-cert.pem \
    -cert $CERT_DIR/server-cert.pem \
    -key $CERT_DIR/server-key.pem \
    $ADDRESS \
    describe $SERVICE_NAME

echo

echo "________________________________________________________________________"

# Mapear o nome do serviço para o arquivo .proto correspondente
# Assumindo que o arquivo .proto tem o mesmo nome que o pacote do serviço
# Exemplo: para `info.Info`, o arquivo seria `info.proto`

# Extrair o pacote e o nome do arquivo .proto
PROTO_FILE="${PROTO_DIR}/$(echo "$SERVICE_NAME" | awk -F. '{print $1}').proto"

if [ -f "$PROTO_FILE" ]; then
    cat "$PROTO_FILE"
else
    echo "Arquivo .proto não encontrado para o serviço $SERVICE_NAME em $PROTO_FILE"
fi

echo "________________________________________________________________________"
