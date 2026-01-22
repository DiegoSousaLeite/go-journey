package main

import (
	"errors"
	"fmt"
)

// DOMÍNIO: Regra de negócio pura.
type Calculadora struct {
	// -------------------------------------------------------------------------
    // CONCEITO DE VISIBILIDADE (ENCAPSULAMENTO) EM GO
    // -------------------------------------------------------------------------
    // Em Java: public double operando1;
    // Em Go:   Operando1 float64 (Começa com Maiúscula = Exportado/Público)
    //
    // Se fosse: operando1 float64 (Começa com minúscula = Não-Exportado/Privado ao pacote)
    //
    // DECISÃO DE DESIGN:
    // Se deixarmos Maiúsculo (Exportado):
    // - Qualquer outro pacote pode fazer: calc.Operando1 = 10.5
    // - É útil para DTOs (Data Transfer Objects) ou structs de configuração.
    //
    // Se mudarmos para minúsculo (Não-Exportado):
    // - Apenas código dentro do "package main" consegue acessar.
    // - Protege o dado de modificações externas acidentais.
    // - Obriga quem for de fora a usar métodos (Setters/Construtores) para definir valores.
    // -------------------------------------------------------------------------
	Operando1 float64
	Operando2 float64
}

func (c Calculadora) Somar() float64 {
	return c.Operando1 + c.Operando2
}

func (c Calculadora) Subtrair() float64 {
	return c.Operando1 - c.Operando2
}

func (c Calculadora) Multiplicar() float64 {
	return c.Operando1 * c.Operando2
}

func (c Calculadora) Dividir() (float64, error) {
	if c.Operando2 == 0 {
		return 0, fmt.Errorf("Divisão por zero não é permitida")
	}
	return c.Operando1 / c.Operando2, nil
}

// CONTROLADOR: O Runner gerencia o fluxo.
// Ele solicita entradas, executa operações e exibe resultados.
type Runner struct {
	// COMPOSIÇÃO: Em vez de herança, o Runner "tem uma" Calculadora.
	//Referência para a struct Calculadora
	calculadora *Calculadora

	// ESTADO: O Runner está guardando o estado da operação atual e do resultado.
	resultado float64
	operacao string
}

// CONSTRUTOR (Idiomático em Go):
// Padrão Factory: "NewNomeDaStruct".
// Recebe a dependência (c *Calculadora) e retorna o ponteiro para o Runner montado.
func NewRunner(c *Calculadora) *Runner {
	return &Runner{calculadora: c}
}

// Solicita os valores dos operandos ao usuário
//Passa o ponteiro do Runner para modificar seu estado
// Se fosse (r Runner), Go faria uma cópia da struct, e nada seria salvo no original.
func (r *Runner) SolicitarValores() error {
	fmt.Println("Digite o primeiro número para iniciar:")

	// Retornando erro se a leitura falhar
	// Armazenando o valor no campo Operando1 da struct Calculadora
	// Em C++ precisaria da seta (r->calculadora->Operando1), 
    // mas Go faz o dereference automático do ponteiro com ponto (.).
	if _, err := fmt.Scanln(&r.calculadora.Operando1); err != nil {
		return errors.New("Erro ao ler o primeiro número")
	}
	// Retornando erro se a leitura falhar
	// Armazenando o valor no campo Operando2 da struct Calculadora
	fmt.Println("Digite o segundo número:")
	if _, err := fmt.Scanln(&r.calculadora.Operando2); err != nil {
		return errors.New("Erro ao ler o segundo número")
	}

	return nil
}

func (r *Runner) SolicitarOperacao() (error) {
	var operacao string
	fmt.Println("Digite a operação desejada (+, -, *, /):")
	/*
	Primeiro, ele executa a parte antes do ; (_, err := ...).
	Depois, ele verifica a condição após o ; (err != nil).
	Se a condição for verdadeira, entra no bloco.
	*/
	// 1. Tente ler e coloque o erro numa variável temporária 'err'
	// 2. IMEDIATAMENTE verifique se 'err' é diferente de nulo (null)
	if _, err := fmt.Scanln(&operacao); err != nil {
		return errors.New("Erro ao ler a operação")
	}

	switch operacao { 
	case "+", "-", "*", "/":
		r.operacao = operacao
		return nil
	default:
		return errors.New("Operação inválida")
	}
}

func (r *Runner) ExecutarOperacao() {
	switch r.operacao {
	case "+":
		r.resultado = r.calculadora.Somar()
	case "-":
		r.resultado = r.calculadora.Subtrair()
	case "*":
		r.resultado = r.calculadora.Multiplicar()
	case "/":
		resultado, err := r.calculadora.Dividir()
		if err != nil {
			fmt.Println("Erro:", err)
		}
		r.resultado = resultado
	}

	fmt.Printf("Resultado: %.2f\n", r.resultado)

}

// Método principal que orquestra o loop infinito
func (r *Runner) Execute() {
	// Go não tem "while". O "for" sem parâmetros é o nosso "while(true)".
	for {
		if err := r.SolicitarValores();  err != nil {
			fmt.Println("Erro:", err)
			continue // Pula para a próxima iteração do loop
		}

		err := r.SolicitarOperacao()
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}
		
		r.ExecutarOperacao()
	}
		
}



func main() {
	// 1. Instanciação da dependência (Calculadora)
    calculadora := &Calculadora{}
    
    // 2. Injeção da dependência no Controlador (Runner) via "Construtor"
    runner := NewRunner(calculadora)
    
    // 3. Disparo da aplicação
    runner.Execute()
}