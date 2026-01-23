package main
import "fmt"

func main(){

	// Declaração e inicialização de um mapa
	//Em maps sempre usamos make para criar um mapa vazio
	mapa := make(map[string]int)

	mapa["Alice"] = 25
	mapa["Bob"] = 30

	fmt.Println("Mapa inicial:", mapa) //output: map[Alice:25 Bob:30 Charlie:35]
}
