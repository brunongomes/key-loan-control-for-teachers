package main

import (
	"log"
	"github.com/gorilla/mux"
	"./routes"
	"net/http"
)

func main() {
	router := mux.NewRouter()	

	// Roteamento para arquivos estáticos
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Chamar a função que configura as rotas
	routes.SetupRoutes(router)

	log.Fatal(http.ListenAndServe(":8081", router))
}
