package main

import (
	"fmt"
	"io"
	"os"
)

func EscrevaOlaMundo(writer io.Writer) (int, error) {
	data := []byte("Hello, Go IO!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return 0, err
	}

	return n, nil
}


func main() {

	//Criar ou abrir um arquivo para escrita
	file, err := os.Create("example.txt")

	if err != nil {
		fmt.Println("Erro ao criar arquivos:", err)
	}

	// Fechar o arquivo ao final da função
	//Defer garante que o arquivo será fechado quando a função main terminar
	//O comando defer em Go agenda uma chamada de função (a "função deferida") para ser executada imediatamente antes de a função que a contém retornar.
	defer file.Close()

	//Escrever "Hello, Go IO!" no arquivo usando a função EscrevaOlaMundo
	n, err := EscrevaOlaMundo(file)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Printf("Escreveu %d bytes no arquivo.\n", n)

}
