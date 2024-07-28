#!/bin/bash

# Endereço do servidor gRPC
ADDRESS="localhost:50051"

# Caminho para os certificados e chaves (se necessário, ajuste conforme seu ambiente)
CERT_DIR="./storage/certs"

# Comando grpcurl para listar todos os serviços e métodos disponíveis usando Reflection
grpcurl -v \
    -cacert $CERT_DIR/ca-cert.pem \
    -cert $CERT_DIR/server-cert.pem \
    -key $CERT_DIR/server-key.pem \
    $ADDRESS \
    list
