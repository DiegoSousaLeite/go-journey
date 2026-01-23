package main
import "fmt"

func main(){

	// Criando um slice usando make
	// make(tipo, tamanho, capacidade)
	// Tem como usar o make apenas com dois parâmetros (tipo e tamanho),
	// nesse caso a capacidade será igual ao tamanho
	slice := make([]int, 0, 5) //Len 0, Cap 5
	fmt.Printf("Slice inicial: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	// Cresncedo dentro da capacidade inicial
	for i := 1; i <= 5; i++ {
		slice = append(slice, i*10)
		fmt.Printf("Após append %d: Slice: %v, Len: %d, Cap: %d\n", i*10, slice, len(slice), cap(slice))
	}

	// Adicionando além da capacidade inicial para ver o crescimento automático
	slice = append(slice, 6)
	fmt.Printf("Após append 60: Slice: %v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

}
