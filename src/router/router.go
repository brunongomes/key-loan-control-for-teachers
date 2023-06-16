package main

import (
	"log"
	"net/http"

	"../internal/disciplinas"
	"../internal/professores"
	"../internal/emprestimos"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Cria um novo roteador
	router := mux.NewRouter()

	// Define as rotas e as funções de manipulador
	router.HandleFunc("/disciplinas", disciplinas.CadastrarDisciplina).Methods("POST")
	router.HandleFunc("/disciplinas", disciplinas.ListarDisciplinas).Methods("GET")
	router.HandleFunc("/disciplinas/{codigo}", disciplinas.ExcluirDisciplina).Methods("DELETE")
	router.HandleFunc("/disciplinas/{codigo}", disciplinas.AtualizarDisciplina).Methods("PUT")

	router.HandleFunc("/professores", professores.CadastrarProfessor).Methods("POST")
	router.HandleFunc("/professores", professores.ListarProfessores).Methods("GET")
	router.HandleFunc("/professores/{cpf}", professores.ExcluirProfessor).Methods("DELETE")
	router.HandleFunc("/professores/{cpf}", professores.AtualizarProfessor).Methods("PUT")

	router.HandleFunc("/emprestimos", emprestimos.CadastrarEmprestimo).Methods("POST")
	router.HandleFunc("/emprestimos", emprestimos.ListarEmprestimos).Methods("GET")
	router.HandleFunc("/emprestimos/{codigo}", emprestimos.ExcluirEmprestimo).Methods("DELETE")
	router.HandleFunc("/emprestimos/{codigo}", emprestimos.AtualizarEmprestimo).Methods("PUT")

	// Inicia o servidor na porta 8080
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8081"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
