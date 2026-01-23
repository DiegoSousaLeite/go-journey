package main
import "fmt"

func main(){
	// Declaração e inicialização de um slice
	// Um slice é mais flexível que um array
	slice := []int{10, 20, 30}
	fmt.Println(slice)
	slice = append(slice, 40, 50)
	fmt.Println(slice)

}
