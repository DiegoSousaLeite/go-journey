package main

import "fmt"

//Exemplo de loop for em Go
func Main1() {

	for i := 1; i <= 10; i++ {
		fmt.Printf("Iteration number: %d\n", i)
	}

}

//Exemplo de função com múltiplos valores de retorno
func dobrar(numero int) (int, int, error) {
	
	if numero < 0 {
		return 0, 0, fmt.Errorf("Negative numbers: %d cannot be doubled", numero)
	}

	
	return numero * 2, numero * 3, nil
}

func main() {
	//No padrão o ultimo retorno é o erro
	//_ é usado para ignorar valores que não serão utilizados
	result, _, err := dobrar(5)
	
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Doubled value: %d\n", result)
	}

}