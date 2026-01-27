package main

import (
	"fmt"
	"io"
	"net/http"
)

func EscrevaOlaMundo(writer io.Writer) (int, error) {
	data := []byte("Hello, Go IO!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return 0, err
	}

	return n, nil
}

func main() {
	//Responde a qualquer verbo HTTP na rota /hello
	//HTTP basico
	http.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().
				Set("Content-Type", "application/json")
			EscrevaOlaMundo(w)
		},
	)
	fmt.Println("Servidor iniciado na porta 8080")
	// Iniciar o servidor HTTP na porta 8080
	http.ListenAndServe(":8080", nil)
}
