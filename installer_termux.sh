#!/bin/bash

if ! command -v go &> /dev/null; then
    echo "Golang não instalado. Instalando..."
    pkg install golang
fi
go build -o hall
cp hall /data/data/com.termux/files/usr/bin
clear
echo 'Hall foi instalado com sucesso. Execute com o comando "hall"'