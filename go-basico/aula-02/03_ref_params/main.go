package main

import "fmt"

func modificarValor(valor int) {
	fmt.Printf("Inside function before change: %d\n", valor)
	valor = 20
	fmt.Printf("Inside function: %d\n", valor)
}

func main() {

	numero := 10
	fmt.Printf("Before function call: %d\n", numero)
	
	//Todos os parâmetros em Go são passados por cópia e não por referência
	modificarValor(numero)
	
	fmt.Printf("After function call: %d\n", numero)
}
