package middleware

//Camada de middleware - Intercepta requisições e respostas HTTP
//Normalmente usada para funcionalidades transversais como logging, autenticação, etc.

import (
	"fmt"
	"net/http"
)

//LoggingMiddleware registra o método e a rota de cada requisição recebida
func LoggingMiddleware(next http.Handler) http.Handler {
	// Retorna um handler que envolve o próximo handler
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf(
				"Método: %s, Rota: %s\n",
				r.Method,
				r.URL.Path,
			)
			next.ServeHTTP(w, r)
		},
	)
}
