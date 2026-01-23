package main

import "fmt"

/*
================================================================================
GUIA DE ARQUITETURA DE DADOS: GO vs JAVA/C++
================================================================================

1. HERANÇA (Inheritance) - "É Um" (Is-A)
   - Java/C++: `class CarroEletrico extends Carro`.
   - Go: NÃO EXISTE palavra-chave `extends`. Go não suporta herança de classes.
   - Alternativa: Usamos Composição (Embedding) para reutilizar código.

2. COMPOSIÇÃO / EMBEDDING (O exemplo deste código) - "É Feito De" (Has-A Forte)
   - Sintaxe: `type CarroEletrico struct { Carro }` (Sem nome de campo, apenas o tipo).
   - Memória: A struct `Carro` é alocada DENTRO do bloco de memória de `CarroEletrico`.
   - Ciclo de Vida: Dependentes. Se o CarroEletrico for destruído, o Carro interno também é.
   - Comportamento: "Method Promotion". O Go faz os métodos de `Carro` (Acelerar/Frear)
     ficarem acessíveis diretamente em `CarroEletrico` (ce.Acelerar), simulando herança.
     OBS: A struct externa (`CarroEletrico`) pode fazer "Shadowing" (sobrescrever) métodos.

3. AGREGAÇÃO - "Tem Um" (Has-A Fraco)
   - Sintaxe: Usaria um PONTEIRO explícito. Ex: `Motor *Motor`.
   - Memória: Blocos separados na memória (Heap), ligados por um endereço (ptr).
   - Ciclo de Vida: Independentes. O `Motor` pode existir sem o `Carro`.
   - Uso Ideal: Quando objetos podem ser trocados (ex: trocar o motorista) ou compartilhados.

4. INTERFACES (Duck Typing)
   - Java: `implements Acelerador`. Explícito.
   - Go: Implícito. Se `Carro` tem o método `Acelerar()`, ele AUTOMATICAMENTE
     implementa a interface `Acelerador`. Não precisa declarar nada.
================================================================================
*/

type Acelerador interface {
	Acelerar()
}

type Freio interface {
	Frear()
}

type Carro struct {
	Modelo string
}

func (c Carro) Acelerar() {
	fmt.Printf("O carro %s está acelerando!\n", c.Modelo)
}

func (c Carro) Frear() {
	fmt.Printf("O carro %s está freando!\n", c.Modelo)
}

// CarroEletrico "herda" os métodos de Carro através da composição
type CarroEletrico struct {
	Carro // Composição: CarroEletrico "tem um" Carro
	BateriaNivel int
}

func (ce CarroEletrico) CarregarBateria() {
	fmt.Printf("Carregando a bateria do carro elétrico %s. Nível atual: %d%%\n", ce.Modelo, ce.BateriaNivel)
}

func main() {
    ce := CarroEletrico{
		Carro: Carro{Modelo: "Tesla Model S"},
		BateriaNivel: 80,
	}

	ce.Acelerar()		// Método herdado de Carro por composição
	ce.Frear()          // Método herdado de Carro por composição
	ce.CarregarBateria() // Método específico de CarroEletrico por composição
}
