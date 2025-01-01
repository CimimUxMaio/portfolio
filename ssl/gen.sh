#!/bin/bash

set -e  # Exit on error

rm -f *.pem
echo "Removed all existing .pem files."

echo "Generating CA's private key and self-signed certificate..."
openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem
echo "ca-key.pem and ca-cert.pem generated."

echo "Generating server's private key and certificate signing request (CSR)..."
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem
echo "server-key.pem and server-req.pem generated."

echo "Signing server's CSR with CA's private key..."
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem
echo "server-cert.pem generated."

echo "Verifying server's certificate..."
openssl verify -CAfile ca-cert.pem server-cert.pem
