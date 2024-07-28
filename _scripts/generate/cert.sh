#!/bin/bash

# Diretório onde serão salvos os certificados
CERT_DIR="./storage/certs"

# Nome do arquivo do certificado e da chave
CERT_FILE="server-cert.pem"
KEY_FILE="server-key.pem"
CA_CERT_FILE="ca-cert.pem"
CA_KEY_FILE="ca-key.pem"

# CN (Common Name) para o certificado
CN="my-server"

# Validade do certificado em dias
VALID_DAYS=365

# Certificado e chave serão gerados em $CERT_DIR
mkdir -p "$CERT_DIR"

# Gerar chave privada da CA
openssl genpkey -algorithm RSA -out "$CERT_DIR/$CA_KEY_FILE"

# Gerar certificado autoassinado da CA
openssl req -x509 -new -nodes -key "$CERT_DIR/$CA_KEY_FILE" -sha256 -days $VALID_DAYS -out "$CERT_DIR/$CA_CERT_FILE" -subj "/CN=MyCA"

# Gerar chave privada do servidor
openssl genpkey -algorithm RSA -out "$CERT_DIR/$KEY_FILE"

# Gerar CSR (Certificate Signing Request) para o servidor com SAN
openssl req -new \
    -key "$CERT_DIR/$KEY_FILE" \
    -out "$CERT_DIR/server-cert.csr" \
    -subj "/CN=$CN" \
    -config <(printf "[req]\ndistinguished_name=req\nreq_extensions=req_ext\n[req_ext]\nsubjectAltName=DNS:localhost")

# Assinar o certificado usando a chave privada da CA com SAN
openssl x509 -req \
    -in "$CERT_DIR/server-cert.csr" \
    -CA "$CERT_DIR/$CA_CERT_FILE" \
    -CAkey "$CERT_DIR/$CA_KEY_FILE" \
    -CAcreateserial \
    -out "$CERT_DIR/$CERT_FILE" \
    -days $VALID_DAYS \
    -extensions req_ext \
    -extfile <(printf "[req_ext]\nsubjectAltName=DNS:localhost")

# Remover CSR após a geração do certificado (opcional)
rm "$CERT_DIR/server-cert.csr"

echo "Certificado e chave gerados em $CERT_DIR"
