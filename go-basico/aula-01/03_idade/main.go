package main

import "fmt"

func main() {

	var age int
	fmt.Println("Please enter your age:")
	fmt.Scanln(&age)

	//Não precisa de () em estruturas de controle
	if age < 18 {
		fmt.Println("You are a minor.")
	} else {
		fmt.Println("You are an adult.")
	}



}