package main

import (
	"fmt"
	"io"
	"os"
)

func LerReader(reader io.Reader, buffer *[]byte) (int, error) {
	n, err := reader.Read(*buffer)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return 0, err
	}	
	return n, nil
}

func main() {

	file, err := os.Open("../01_iowriter/example.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}

	defer file.Close()

	buffer := make([]byte, 1024) // Criar um buffer para armazenar os dados lidos

	n, err := LerReader(file, &buffer)
	if err != nil {
		fmt.Println("Erro ao ler buffer:", err)
		return
	}

	fmt.Printf("Leu %d bytes do arquivo:\n", n)
	fmt.Println(string(buffer[:n]))


}
