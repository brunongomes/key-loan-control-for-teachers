package routes

import (
    "../internal/disciplinas"
    "../internal/professores"
    "../internal/emprestimos"
    "github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
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
}
