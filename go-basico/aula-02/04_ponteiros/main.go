package main
import "fmt"

/*
& - Operador de endereço: Retorna o endereço de memória de uma variável.
* - Operador de desreferenciação: Acessa o valor armazenado no endereço apontado por um ponteiro.

Existe algo que o GO não faz por segurança: Aritmética de ponteiros. Você não pode incrementar ou decrementar ponteiros como em outras linguagens.

Ponteiros são úteis para:
1. Passagem eficiente de grandes estruturas para funções, evitando cópias desnecessárias.
2. Modificação de variáveis em funções, permitindo que a função altere o valor original.
3. Implementação de estruturas de dados dinâmicas, como listas ligadas e árvores.
*/

func Main1() {

	//Uma varialvel normal
	idade := 30
	//Um ponteiro que armazena o endereço da variável idade
	ponteiroIdade := &idade
	
	fmt.Printf("Valor da idade: %d\n", idade)
	//Imprimindo o endereço da variável idade
	fmt.Printf("Endereço da variável idade: %p\n", &idade)
	//Imprimindo o valor do ponteiro (endereço da variável idade)
	fmt.Printf("Valor do ponteiro ponteiroIdade: %p\n", ponteiroIdade)
	//Imprimindo o valor apontado pelo ponteiro (valor da variável idade)
	fmt.Printf("Valor apontado pelo ponteiro ponteiroIdade: %d\n", *ponteiroIdade)

}

func main() {
	numero := 42
	ponteiroNumero := &numero

	//Mostrando valor antes da modificação via ponteiro
	fmt.Println("Before: ", numero)
	//Modificando o valor da variável numero via ponteiro
	*ponteiroNumero = 100
	//Mostrando valor após a modificação via ponteiro
	fmt.Println("After: ", numero)
}