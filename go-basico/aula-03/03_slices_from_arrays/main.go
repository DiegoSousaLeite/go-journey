package main
import "fmt"

func main(){

	// Declaração e inicialização de um array
	array := [5]int{10, 20, 30, 40, 50}

	// Criando um slice a partir de um array
	slice := array[1:4] // Inclui os elementos do índice 1 ao 3 (4 é exclusivo)

	//%v é usado para formatar valores de forma padrão
	//Ele é utilizado para imprimir valores de qualquer tipo em Go
	fmt.Printf("Array: %v\n", array)

	// Imprimindo o slice, seu tamanho (len) e capacidade (cap)
	//cap é a capacidade máxima do slice antes de precisar alocar mais memória
	//Aqui é 4 porque o slice começa no índice 1 do array original que tem 5 elementos
	//Capacidade = (Tamanho Total do Array) - (Índice de Início do Slice)
	fmt.Printf("Slice: %v, Len: %d, Cap: %d", slice, len(slice), cap(slice))

	slice[0] = 100
	
	fmt.Printf("\nArray após modificação do slice: %v\n", array)
	fmt.Printf("Slice após modificação: %v\n", slice)

}
