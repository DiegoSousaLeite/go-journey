#!/bin/bash

# Verifica se o usuário passou o caminho como argumento
if [ -z "$1" ]; then
    echo "Uso: ./criar_projeto.sh <caminho/para/o/projeto>"
    exit 1
fi

# O caminho desejado (ex: go-basico/aula-03/07_map)
CAMINHO="$1"

# 1. Cria a estrutura de pastas (se já existirem, o -p não faz nada e não dá erro)
mkdir -p "$CAMINHO"

# Define o caminho do arquivo main.go
ARQUIVO="$CAMINHO/main.go"

# 2. Verifica se o arquivo main.go JÁ existe
if [ -f "$ARQUIVO" ]; then
    echo "Aviso: O arquivo '$ARQUIVO' já existe. Nenhuma alteração foi feita."
else
    # Se não existe, cria o arquivo com o conteúdo boilerplate
    cat <<EOF > "$ARQUIVO"
package main

import "fmt"

func main() {
    
}
EOF
    echo "Sucesso! Projeto criado em: $CAMINHO"
fi