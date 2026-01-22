package main
import "fmt"

// Estrutura da Calculadora
type Calculadora struct {
	Operando1 float64
	Operando2 float64
}

func (c Calculadora) Dividir() (float64, error) {
	if c.Operando2 == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return c.Operando1 / c.Operando2, nil
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


func main() {
	var op1, op2 float64
	var operacao string
	// Solicitando entrada do usuário
	fmt.Println("Digite a operação (+, -, *, /):")
	fmt.Scanln(&operacao)
	fmt.Println("Digite o primeiro operando:")
	fmt.Scanln(&op1)
	fmt.Println("Digite o segundo operando:")
	fmt.Scanln(&op2)
	calculadora := Calculadora{op1, op2}
	// Realizando a operação com base na entrada do usuário
	// Usando switch para selecionar a operação
	switch operacao {
	case "+":
		resultado := calculadora.Somar()
		fmt.Printf("Resultado: %.2f\n", resultado)
	case "-":
		resultado := calculadora.Subtrair()
		fmt.Printf("Resultado: %.2f\n", resultado)
	case "*":
		resultado := calculadora.Multiplicar()
		fmt.Printf("Resultado: %.2f\n", resultado)
	case "/":
		resultado, err := calculadora.Dividir()
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Printf("Resultado: %.2f\n", resultado)
		}
	default:
		fmt.Println("Operação inválida")
	}
}