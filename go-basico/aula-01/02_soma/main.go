package main

import "fmt"

func sum(n1, n2 int) int {
	return n1 + n2
}

func main() {

	var a, b int
	fmt.Println("Enter two integers to sum:")
	fmt.Scanln(&a, &b)
	//:=(Goffer) é um atalho para declarar e inicializar uma variável 
	sum := sum(a, b)
	fmt.Printf("The sum of %d and %d is %d.\n", a, b, sum)


}