package main

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"bytes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Disciplina struct {
	Codigo       string `json:"codigo"`
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"cargaHoraria"`
}

func main() {

	router := mux.NewRouter()
	// Roteamento para arquivos estáticos
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/disciplinas", disciplinasHandler).Methods("GET")
	router.HandleFunc("/emprestimos", emprestimosHandler).Methods("GET")
	router.HandleFunc("/professores", professoresHandler).Methods("GET")
	router.HandleFunc("/disciplinas", cadastrarDisciplinaHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Home</title>
				<link rel="stylesheet" href="/static/style.css">
			</head>
			<body>
				<h1>Home</h1>
				<ul>
					<li><a href="/disciplinas">Disciplinas</a></li>
					<li><a href="/emprestimos">Empréstimos</a></li>
					<li><a href="/professores">Professores</a></li>
				</ul>
			</body>
			</html>
			`
		fmt.Fprintf(w, html)
	}
}

func disciplinasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Exibir a página de disciplinas
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Cadastrar Disciplina</title>
		</head>
		<body>
			<h1>Cadastrar Disciplina</h1>
			<form action="/disciplinas" method="POST">
				<label for="codigo">Código:</label><br>
				<input type="text" id="codigo" name="codigo"><br><br>
				<label for="nome">Nome:</label><br>
				<input type="text" id="nome" name="nome"><br><br>
				<label for="cargaHoraria">Carga Horária:</label><br>
				<input type="text" id="cargaHoraria" name="cargaHoraria"><br><br>
				<input type="submit" value="Cadastrar">
			</form>
		</body>
		</html>
			`
		fmt.Fprintf(w, html)
	}
}

func emprestimosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Exibir a página de empréstimos
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Cadastrar Disciplina</title>
		</head>
		<body>
			<h1>Cadastrar Empréstimos</h1>
			<form action="/" method="POST">
				<label for="codigo">Código:</label><br>
				<input type="text" id="codigo" name="codigo"><br><br>
				<label for="cpfProfessor">CPF do professor:</label><br>
				<input type="text" id="cpfProfessor" name="cpfProfessor"><br><br>
				<label for="nomeProfessor">Nome do professor:</label><br>
				<input type="text" id="nomeProfessor" name="nomeProfessor"><br><br>
				<label for="horarioInicio">Horário início:</label><br>
				<input type="text" id="horarioInicio" name="horarioInicio"><br><br>
				<label for="horarioFim">Horário início:</label><br>
				<input type="text" id="horarioFim" name="horarioFim"><br><br>
				<input type="submit" value="Cadastrar">
			</form>
		</body>
		</html>
			`
		fmt.Fprintf(w, html)
	}
}

func professoresHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Exibir a página de professores
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Cadastrar Professor</title>
		</head>
		<body>
			<h1>Cadastrar Professor</h1>
			<form action="/" method="POST">
				<label for="cpf">CPF:</label><br>
				<input type="text" id="cpf" name="cpf"><br><br>
				<label for="nome">Nome:</label><br>
				<input type="text" id="nome" name="nome"><br><br>
				<input type="submit" value="Cadastrar">
			</form>
		</body>
		</html>
			`
		fmt.Fprintf(w, html)
	}
}

func cadastrarDisciplinaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
			// Obter os dados do formulário
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Erro ao analisar os dados do formulário", http.StatusBadRequest)
				return
			}

			codigo := strings.TrimSpace(r.Form.Get("codigo"))
			nome := strings.TrimSpace(r.Form.Get("nome"))
			cargaHorariaStr := strings.TrimSpace(r.Form.Get("cargaHoraria"))

			// Validar os dados do formulário (opcional)

			// Converter a carga horária para o tipo int
			cargaHoraria, err := strconv.Atoi(cargaHorariaStr)
			if err != nil {
				http.Error(w, "Carga horária inválida", http.StatusBadRequest)
				return
			}

			// Criar a estrutura de dados para a disciplina
			disciplina := Disciplina{
				Codigo:       codigo,
				Nome:         nome,
				CargaHoraria: cargaHoraria,
			}

			// Converter a estrutura de dados para JSON
			payload, err := json.Marshal(disciplina)
			if err != nil {
				http.Error(w, "Erro ao converter os dados para JSON", http.StatusInternalServerError)
				log.Println("Erro ao converter os dados para JSON:", err) // Adiciona um log de erro
				return
			}

			// Enviar a solicitação HTTP POST para o backend
			resp, err := http.Post("http://localhost:8080/disciplinas", "application/json", bytes.NewBuffer(payload))
			if err != nil {
				http.Error(w, "Erro ao enviar a solicitação HTTP para o backend", http.StatusInternalServerError)
				log.Println("Erro ao enviar a solicitação HTTP para o backend:", err) // Adiciona um log de erro
				return
			}
			defer resp.Body.Close()

			// Verificar a resposta do backend
			if resp.StatusCode == http.StatusCreated {
				fmt.Fprintf(w, "Disciplina cadastrada com sucesso")
			} else {
				http.Error(w, "Erro ao cadastrar a disciplina", http.StatusInternalServerError)
				log.Println("Erro ao cadastrar a disciplina: status", resp.StatusCode) // Adiciona um log de erro
			}
	}
}
