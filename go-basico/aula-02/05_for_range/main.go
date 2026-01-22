package main

import "fmt"

func main() {

	numeros := []int{10, 20, 30, 40, 50}

	// Já que não é possivel fazer aritmética de ponteiros em Go,
	// o range retorna uma cópia do valor em cada iteração.
	// Portanto, o endereço impresso será sempre o do valor temporário criado na iteração.
	// Se quiser o endereço do elemento original, é necessário usar o índice para acessá-lo no slice.
	for indice, numero := range numeros {
		fmt.Printf("Índice: %d, Numero: %d, Endereço: %p\n", indice, numero, &numero)
	}

}