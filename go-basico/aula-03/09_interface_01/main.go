package main

import "fmt"

type Carro struct {
	Modelo string 
}

func (c Carro) Acelerar() {
	fmt.Printf("O carro %s está acelerando!\n", c.Modelo)
}

type Bicicleta struct {
	Tipo string
}

func (b Bicicleta) Acelerar() {
	fmt.Printf("A bicicleta do tipo %s está acelerando!\n", b.Tipo)
}


// Interfaces são sempre abstratos/sempre ponteiros em Go
// Quem se importa com a interface é quem ta usando ela, quem está precisando do metodo
type Veiculo interface {
	Acelerar()
}

func TestaVeiculo(v Veiculo) {
	v.Acelerar()
}

func main() {
	meuCarro := Carro{Modelo: "Toyota"}
	minhaBicicleta := Bicicleta{Tipo: "Mountain Bike"}

	TestaVeiculo(meuCarro)
	TestaVeiculo(minhaBicicleta)
    
}
