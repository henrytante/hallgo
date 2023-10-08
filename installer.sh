#!/bin/bash

if [ "$EUID" -ne 0 ]; then
  echo "Execute o script como root"
  exit 1
fi

if ! command -v go &> /dev/null; then
    echo "Golang não instalado. Instalando..."
    sudo apt install golang
fi
go build -o hall
sudo cp hall /bin
clear
echo 'Hall foi instalado com sucesso. Execute com o comando "hall"'