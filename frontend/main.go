package main

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
)

func main() {
	// Define a rota raiz para servir o arquivo index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/home/bruno/Área de Trabalho/sistema-emprestimo-chaves/front/index.html")
	})

	// Middleware para adicionar o cabeçalho Access-Control-Allow-Origin
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
			next.ServeHTTP(w, r)
		})
	}

	// Inicia o servidor na porta 8081 com o tratamento do CORS
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:8081"}))(corsMiddleware(http.DefaultServeMux))))
}
