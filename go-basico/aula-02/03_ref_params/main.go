package main

import "fmt"

// Função que modifica o valor original usando ponteiro
func mofificarOriginal(valor *int) {
	*valor = *valor * 2
}

func main() {

	numero := 10
	fmt.Printf("Before function call: %d\n", numero)
	
	// Passando o endereço da variável numero
	mofificarOriginal(&numero)
	
	fmt.Printf("After function call: %d\n", numero)
}
