package main
import "fmt"

func main(){
	 
	// Declaração de um slice vazio
	var slice []int
	
	fmt.Printf("Before append: Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	//Crescimento do slice e observação de len e cap
	//Quando a capacidade atual é excedida, o Go aloca um novo array maior, copia os elementos antigos e adiciona o novo elemento
	//A capacidade geralmente dobra quando é necessário mais espaço
	for i := 1; i <= 9; i++ {
		slice = append(slice, i)
		fmt.Printf("Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))
	}
}