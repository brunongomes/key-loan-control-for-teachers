package main

import (
    "log"
    "net/http"
		"./routes"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

func main() {
    // Cria um novo roteador
    router := mux.NewRouter()

    // Chama a função que configura as rotas
    routes.SetupRoutes(router)

    // Inicia o servidor na porta 8080
    allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8081"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
