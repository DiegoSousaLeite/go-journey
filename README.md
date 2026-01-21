# Go Journey 🚀

Este repositório documenta minha trajetória de aprendizado na linguagem **Go (Golang)**, migrando de um background em Java e C++. O foco aqui é entender não apenas a sintaxe, mas a filosofia da linguagem, gerenciamento de memória e arquitetura de microsserviços.

O projeto segue a trilha educacional do **Lucas Badico (Go Básico)**, evoluindo de scripts simples para APIs completas e gRPC.

## 📚 Conteúdo Estudado

### 🟢 Aula 01: Fundamentos e Filosofia
O primeiro passo focou em entender como o Go resolve problemas de build lento e complexidade excessiva de outras linguagens.
- **Ambiente & Tooling:** Configuração do workspace, `go mod`, `go build` vs `go run`.
- **Sintaxe Core:** Declaração de variáveis (inferência com `:=` vs `var`), tipos primitivos (`int`, `float64`, `bool`).
- **Controle de Fluxo:** O uso do `for` como laço único (substituindo while/do-while) e condicionais.
- **Input/Output:** Interação com terminal usando pacote `fmt` e `Scanln`.
- **Tratamento de Erros:** Introdução à filosofia de erros como valores explícitos.
- **Desafio:** Cálculo de ano de nascimento usando pacote `time`.

### 🟡 Aula 02: Memória, Ponteiros e Structs
Aprofundamento no gerenciamento de memória (essencial para quem vem do C++) e organização de dados.
- **Funções:** Múltiplos retornos e "named returns".
- **Ponteiros:** Diferença entre *Pass-by-value* (padrão do Go) e *Pass-by-reference*. Uso de `&` e `*` para manipulação direta de memória e otimização.
- **Structs:** Modelagem de dados e associação de métodos (substituindo classes).
- **Projeto Prático:** Desenvolvimento de uma **Calculadora Interativa** no terminal utilizando um padrão `Runner` para encapsular a lógica e validar entradas.

### 🟠 Aula 03: Coleções e Design de Software (A "OOP" do Go)
Como o Go aborda Orientação a Objetos através de **Composição ao invés de Herança**.
- **Coleções:** Diferenças críticas entre Arrays (fixos) e Slices (dinâmicos). Manipulação de capacidade (`cap` vs `len`), função `make` e `append`.
- **Maps:** Estruturas chave-valor.
- **Interfaces & Polimorfismo:** Definição de contratos de comportamento. Entendimento de que interfaces definem "o que" fazer, permitindo desacoplamento.
- **Modularidade:** Organização de pacotes (Public vs Private), encapsulamento e `go mod tidy`.
- **Refatoração:** Evolução da calculadora para um design modular utilizando Interfaces para as operações matemáticas.

### 🔵 Aula 04: APIs REST e Interfaces I/O
Transição para o desenvolvimento web e entendimento profundo das interfaces nativas.
- **net/http:** Estrutura de Requests/Responses, Headers e Body.
- **Abstração de I/O:** O poder das interfaces `io.Reader` e `io.Writer` para manipular arquivos e streams HTTP de forma intercambiável.
- **Arquitetura:** Organização de projeto em camadas: *Handlers*, *Services* e *Models*.
- **Middlewares:** Implementação de interceptadores para logs e validações.
- **Frameworks:** Comparativo entre a standard library e frameworks como GorillaMux, Echo e Fiber.

### 🟣 Aula 05: gRPC e Microsserviços
Implementação de comunicação de alta performance entre serviços.
- **HTTP vs gRPC:** Diferenças de performance, serialização e contratos.
- **Protocol Buffers:** Criação de `.proto` files para definição de contratos rigorosos.
- **Code Generation:** Uso do `protoc` para gerar o código Go automaticamente.
- **Server Implementation:** Criação de um serviço de usuários (`CreateUser`, `GetUser`) funcional.
- **Interceptors:** O equivalente a middlewares no ecossistema gRPC.

---

## 🛠 Tech Stack
- **Linguagem:** Go 1.x
- **Comunicação:** REST (net/http), gRPC (Protobuf)
- **Conceitos:** Ponteiros, Interfaces, Concorrência, Clean Architecture.

---
*Estudos realizados por [Diego Sousa Leite](https://github.com/SEU-USUARIO)*