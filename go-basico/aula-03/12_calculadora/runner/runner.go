package runner

import (
	"fmt"
	"github.com/Diego/calculadora/operacao"
)

// As melhor interfaces são aquelas enxutas 
// É aquela que tem no máximo 2 métodos
// Repare que a interface está onde ela é usada, não onde é implementada
type Operacao interface {
	Calcular(a, b float64) float64
}

type Runner struct {
	Operacoes map[string]Operacao
}

func (r *Runner) Executar() {
	var a, b float64
	var operacao string
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&a)
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&b)
	fmt.Println("Digite a operação (+, -, *, /):")
	fmt.Scan(&operacao)

	//r.Operacoes é um map[string]Operacao
	// r.Operacoes[operacao] retorna dois valores: o valor associado à chave "operacao" e um booleano indicando se a chave existe no map
	//op é do tipo Operacao
	//existe é um booleano que indica se a operação foi encontrada no map
	op, existe := r.Operacoes[operacao]
	if !existe {
		fmt.Println("Operação inválida!")
		return
	}

	resultado := op.Calcular(a, b)
	fmt.Printf("Resultado: %.2f\n", resultado)
}

// Construtor do Runner
func NewRunner() *Runner {
	return &Runner{
		// Basicamente um Builder
		Operacoes: map[string]Operacao{
			"+": operacao.Soma{},
			"-": operacao.Subtracao{},
			"*": operacao.Multiplicacao{},
			"/": operacao.Divisao{},
		},
	}
}