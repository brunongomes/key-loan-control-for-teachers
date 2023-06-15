package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"./src/pkg/mongodb"
)

func main() {
	// Cria um novo roteador
	router := mux.NewRouter()

	// Define as rotas e as funções de manipulador
	router.HandleFunc("/disciplinas", CadastrarDisciplina).Methods("POST")
	router.HandleFunc("/disciplinas", ListarDisciplinas).Methods("GET")
	router.HandleFunc("/disciplinas/{codigo}", ExcluirDisciplina).Methods("DELETE")
	router.HandleFunc("/disciplinas/{codigo}", AtualizarDisciplina).Methods("PUT")

	router.HandleFunc("/professores", CadastrarProfessor).Methods("POST")
	router.HandleFunc("/professores", ListarProfessores).Methods("GET")
	router.HandleFunc("/professores/{cpf}", ExcluirProfessor).Methods("DELETE")
	router.HandleFunc("/professores/{cpf}", AtualizarProfessor).Methods("PUT")

	router.HandleFunc("/emprestimos", CadastrarEmprestimo).Methods("POST")
	router.HandleFunc("/emprestimos", ListarEmprestimos).Methods("GET")
	router.HandleFunc("/emprestimos/{codigo}", ExcluirEmprestimo).Methods("DELETE")
	router.HandleFunc("/emprestimos/{codigo}", AtualizarEmprestimo).Methods("PUT")

	// Inicia o servidor na porta 8080
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8081"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

type Disciplina struct {
	Codigo       string `json:"codigo"`
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"cargaHoraria"`
}

func CadastrarDisciplina(w http.ResponseWriter, r *http.Request) {
	// Decodifica o JSON recebido na requisição em uma struct Disciplina
	var disciplina Disciplina
	err := json.NewDecoder(r.Body).Decode(&disciplina)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}

	// Insere a disciplina no banco de dados
	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	err = db.Insert("disciplinas", disciplina)
	if err != nil {
		http.Error(w, "Erro ao cadastrar a disciplina", http.StatusInternalServerError)
		return
	}

	// Configura o cabeçalho Access-Control-Allow-Origin para permitir a origem do frontend
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Disciplina cadastrada com sucesso")
}

// Restante do código omitido por brevidade



func ListarDisciplinas(w http.ResponseWriter, r *http.Request) {
	// Lógica de listagem de disciplinas aqui
	fmt.Fprintf(w, "Listagem de disciplinas")
}

func ExcluirDisciplina(w http.ResponseWriter, r *http.Request) {
	// Lógica de exclusão de disciplina aqui
	fmt.Fprintf(w, "Exclusão de disciplina")
}

func AtualizarDisciplina(w http.ResponseWriter, r *http.Request) {
	// Lógica de atualização de disciplina aqui
	fmt.Fprintf(w, "Atualização de disciplina")
}

func CadastrarProfessor(w http.ResponseWriter, r *http.Request) {
	// Lógica de cadastro de professor aqui
	fmt.Fprintf(w, "Cadastro de professor")
}

func ListarProfessores(w http.ResponseWriter, r *http.Request) {
	// Lógica de listagem de professores aqui
	fmt.Fprintf(w, "Listagem de professores")
}

func ExcluirProfessor(w http.ResponseWriter, r *http.Request) {
	// Lógica de exclusão de professor aqui
	fmt.Fprintf(w, "Exclusão de professor")
}

func AtualizarProfessor(w http.ResponseWriter, r *http.Request) {
	// Lógica de atualização de professor aqui
	fmt.Fprintf(w, "Atualização de professor")
}

func CadastrarEmprestimo(w http.ResponseWriter, r *http.Request) {
	// Lógica de cadastro de empréstimo aqui
	fmt.Fprintf(w, "Cadastro de empréstimo")
}

func ListarEmprestimos(w http.ResponseWriter, r *http.Request) {
	// Lógica de listagem de empréstimos aqui
	fmt.Fprintf(w, "Listagem de empréstimos")
}

func ExcluirEmprestimo(w http.ResponseWriter, r *http.Request) {
	// Lógica de exclusão de empréstimo aqui
	fmt.Fprintf(w, "Exclusão de empréstimo")
}

func AtualizarEmprestimo(w http.ResponseWriter, r *http.Request) {
	// Lógica de atualização de empréstimo aqui
	fmt.Fprintf(w, "Atualização de empréstimo")
}
