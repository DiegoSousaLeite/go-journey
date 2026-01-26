# criar_projeto.ps1

# Define que o script espera um parâmetro (o caminho)
param (
    [Parameter(Mandatory=$true, HelpMessage="Uso: .\criar_projeto.ps1 <caminho/para/o/projeto>")]
    [string]$Caminho
)

# 1. Cria a estrutura de pastas (o -Force age como o -p do linux, não dá erro se já existir)
New-Item -Path $Caminho -ItemType Directory -Force | Out-Null

# Define o caminho do arquivo main.go usando Join-Path (que resolve as barras \ ou / automaticamente)
$Arquivo = Join-Path -Path $Caminho -ChildPath "main.go"

# 2. Verifica se o arquivo main.go JÁ existe
if (Test-Path $Arquivo) {
    Write-Host "Aviso: O arquivo '$Arquivo' já existe. Nenhuma alteração foi feita." -ForegroundColor Yellow
}
else {
    # Se não existe, cria o arquivo com o conteúdo boilerplate
    $Conteudo = @"
package main

import "fmt"

func main() {
    
}
"@
    Set-Content -Path $Arquivo -Value $Conteudo -Encoding UTF8
    Write-Host "Sucesso! Projeto criado em: $Caminho" -ForegroundColor Green
}