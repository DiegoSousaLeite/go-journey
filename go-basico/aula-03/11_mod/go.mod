//comando para criar o módulo:
//go mod init github.com/Diego/meuprimeiromodulo
module github.com/Diego/meuprimeiromodulo

go 1.25.6

//comando para adicionar a biblioteca cobra:
//go get github.com/spf13/cobra
// require (
// 	github.com/inconshreveable/mousetrap v1.1.0 // indirect
// 	github.com/spf13/cobra v1.10.2 // indirect
// 	github.com/spf13/pflag v1.0.9 // indirect
// )

//Adicionar Dependências Faltantes: 
//Ele inspeciona todos os seus arquivos .go, encontra as importações (import) de pacotes que não estão listados no go.mod e os adiciona.
//Remover Dependências Desnecessárias: Ele remove do go.mod e go.sum quaisquer dependências que são listadas, mas que não são mais utilizadas em nenhum lugar do seu código.
// go mod tidy
