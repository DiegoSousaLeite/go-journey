package main

import "fmt"

//Exemplo de loop for em Go
func Main1() {

	for i := 1; i <= 10; i++ {
		fmt.Printf("Iteration number: %d\n", i)
	}

}

//Exemplo de loop infinito com break e continue
func main() {
	var age int
	for {
		fmt.Println("Please enter your age (Must be over 18.):")
		/*
		// A função retorna: (int, error)
		quantidadeLida, erroDetectado := fmt.Scanln(&age)

		// Como não vamos usar "quantidadeLida", o Go reclamaria.
		// Então usamos o "_" para ignorar o primeiro valor.
		_, erroDetectado := fmt.Scanln(&age)
		*/
		_, err := fmt.Scanln(&age)

		//nil em Go é similar a null em outras linguagens
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid age.")
			continue
		}

		if age < 18 {
			fmt.Println("Invalid input. Please enter a valid age.")
			continue
		} 
		break
	}
	fmt.Printf("You have entered a valid age: %d\n", age)

}