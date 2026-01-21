package main

import "fmt"

func main() {

	//Não precisa de ; no final das linhas
	var name string
	fmt.Println("Hello, what is your name?")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s! Welcome to Go programming.\n", name)


}