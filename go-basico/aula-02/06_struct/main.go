package main

import "fmt"

// Definição da struct Pessoa
type Pessoa struct {
	Nome  string
	Maior bool
	Idade int
}

// Método para apresentar a pessoa
//Associado ao tipo Pessoa
func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Nome: %s, Idade: %d, Maior de idade: %t", p.Nome, p.Idade, p.Maior)
}

// Método para envelhecer a pessoa em 1 ano
//Associado ao tipo Pessoa
// Usando ponteiro para modificar o valor original
func (p *Pessoa) Envelhecer() {
	p.Idade++
}

func main() {
	p := Pessoa{"Diego", true, 23}
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Maior de idade:", p.Maior)
	fmt.Println("Idade:", p.Idade)
	fmt.Println("Pessoa completa:", p)
	fmt.Println("Pessoa &", &p)
	fmt.Println("Nome &: ", &p.Nome)
	fmt.Println("Maior &: ", &p.Maior)
	fmt.Println("Idade &: ", &p.Idade)
	fmt.Println(p.Apresentar())
	p.Envelhecer()
	fmt.Println("Após envelhecer:", p.Apresentar())

}