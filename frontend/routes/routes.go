package routes

import (
	"github.com/gorilla/mux"
	"../handlers"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/disciplinas", handlers.DisciplinasHandler).Methods("GET")
	router.HandleFunc("/disciplinas", handlers.CadastrarDisciplinaHandler).Methods("POST")
	router.HandleFunc("/visualizar-disciplinas", handlers.VisualizarDisciplinasHandler).Methods("GET")
	router.HandleFunc("/deletar-disciplina/{codigo}", handlers.DeletarDisciplinaHandler).Methods("DELETE")

	router.HandleFunc("/professores", handlers.ProfessoresHandler).Methods("GET")
	router.HandleFunc("/professores", handlers.CadastrarProfessorHandler).Methods("POST")
	router.HandleFunc("/visualizar-professores", handlers.VisualizarProfessoresHandler).Methods("GET")
	
	router.HandleFunc("/emprestimos", handlers.EmprestimosHandler).Methods("GET")
	router.HandleFunc("/emprestimos", handlers.CadastrarEmprestimoHandler).Methods("POST")
	router.HandleFunc("/visualizar-emprestimos", handlers.VisualizarEmprestimosHandler).Methods("GET")
}