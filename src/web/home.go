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
	"html/template"
)

type Disciplina struct {
	Codigo       string `json:"codigo"`
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"cargaHoraria"`
}

type Emprestimo struct {
	Codigo         int    `json:"codigo"`
	CPF_Professor  string `json:"cpfProfessor"`
	Nome_Professor string `json:"nomeProfessor"`
	Horario_inicio string `json:"horarioInicio"`
	Horario_fim    string `json:"horarioFim"`
}

type Professor struct {
	CPF  string `json:"cpf"`
	Nome string `json:"nome"`
}


func main() {

	router := mux.NewRouter()
	// Roteamento para arquivos estáticos
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/disciplinas", disciplinasHandler).Methods("GET")
	router.HandleFunc("/visualizar-disciplinas", visualizarDisciplinasHandler).Methods("GET") // Nova rota para visualizar disciplinas
	router.HandleFunc("/emprestimos", emprestimosHandler).Methods("GET")
	router.HandleFunc("/professores", professoresHandler).Methods("GET")
	router.HandleFunc("/disciplinas", cadastrarDisciplinaHandler).Methods("POST")
	router.HandleFunc("/professores", cadastrarProfessorHandler).Methods("POST")
	router.HandleFunc("/emprestimos", cadastrarEmprestimoHandler).Methods("POST")

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
			<br>
			<form action="/" method="GET">
				<input type="submit" value="Voltar">
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
			<br>
			<form action="/" method="GET">
				<input type="submit" value="Voltar">
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
			<br>
			<form action="/" method="GET">
				<input type="submit" value="Voltar">
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

func visualizarDisciplinasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Obtenha as disciplinas do banco de dados (substitua esta parte com a lógica de obtenção dos dados do banco)

		// Exiba a página de visualização das disciplinas
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Visualizar Disciplinas</title>
		</head>
		<body>
			<h1>Visualizar Disciplinas</h1>
			<table>
				<thead>
					<tr>
						<th>Código</th>
						<th>Nome</th>
						<th>Carga Horária</th>
						<th>Atualizar</th>
						<th>Deletar</th>
					</tr>
				</thead>
				<tbody>
					{{range .}}
					<tr>
						<td>{{.Codigo}}</td>
						<td>{{.Nome}}</td>
						<td>{{.CargaHoraria}}</td>
						<td><a href="/atualizar-disciplina?codigo={{.Codigo}}">Atualizar</a></td>
						<td><a href="/deletar-disciplina?codigo={{.Codigo}}">Deletar</a></td>
					</tr>
					{{end}}
				</tbody>
			</table>
			<br>
			<form action="/" method="GET">
				<input type="submit" value="Voltar">
			</form>
		</body>
		</html>
		`

		// Substitua esta parte com a lógica para obter as disciplinas do banco
		disciplinas := []Disciplina{
			{Codigo: "D001", Nome: "Disciplina 1", CargaHoraria: 60},
			{Codigo: "D002", Nome: "Disciplina 2", CargaHoraria: 80},
			{Codigo: "D003", Nome: "Disciplina 3", CargaHoraria: 120},
		}

		// Renderize o HTML substituindo os dados das disciplinas
		tmpl, err := template.New("visualizar_disciplinas").Parse(html)
		if err != nil {
			http.Error(w, "Erro ao renderizar o HTML", http.StatusInternalServerError)
			log.Println("Erro ao renderizar o HTML:", err) // Adiciona um log de erro
			return
		}

		err = tmpl.Execute(w, disciplinas)
		if err != nil {
			http.Error(w, "Erro ao renderizar o HTML", http.StatusInternalServerError)
			log.Println("Erro ao renderizar o HTML:", err) // Adiciona um log de erro
			return
		}
	}
}


func cadastrarProfessorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obter os dados do formulário
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erro ao analisar os dados do formulário", http.StatusBadRequest)
			return
		}

		cpf := strings.TrimSpace(r.Form.Get("cpf"))
		nome := strings.TrimSpace(r.Form.Get("nome"))

		// Validar os dados do formulário (opcional)

		// Criar a estrutura de dados para o professor
		professor := Professor{
			CPF:  cpf,
			Nome: nome,
		}

		// Converter a estrutura de dados para JSON
		payload, err := json.Marshal(professor)
		if err != nil {
			http.Error(w, "Erro ao converter os dados para JSON", http.StatusInternalServerError)
			log.Println("Erro ao converter os dados para JSON:", err) // Adiciona um log de erro
			return
		}

		// Enviar a solicitação HTTP POST para o backend
		resp, err := http.Post("http://localhost:8080/professores", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			http.Error(w, "Erro ao enviar a solicitação HTTP para o backend", http.StatusInternalServerError)
			log.Println("Erro ao enviar a solicitação HTTP para o backend:", err) // Adiciona um log de erro
			return
		}
		defer resp.Body.Close()

		// Verificar a resposta do backend
		if resp.StatusCode == http.StatusCreated {
			fmt.Fprintf(w, "Professor cadastrado com sucesso")
		} else {
			http.Error(w, "Erro ao cadastrar o professor", http.StatusInternalServerError)
			log.Println("Erro ao cadastrar o professor: status", resp.StatusCode) // Adiciona um log de erro
		}
	}
}

func cadastrarEmprestimoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obter os dados do formulário
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erro ao analisar os dados do formulário", http.StatusBadRequest)
			return
		}

		codigoStr := strings.TrimSpace(r.Form.Get("codigo"))
		cpfProfessor := strings.TrimSpace(r.Form.Get("cpfProfessor"))
		nomeProfessor := strings.TrimSpace(r.Form.Get("nomeProfessor"))
		horarioInicio := strings.TrimSpace(r.Form.Get("horarioInicio"))
		horarioFim := strings.TrimSpace(r.Form.Get("horarioFim"))

		// Validar os dados do formulário (opcional)

		// Converter o código para o tipo int
		codigo, err := strconv.Atoi(codigoStr)
		if err != nil {
			http.Error(w, "Código inválido", http.StatusBadRequest)
			return
		}

		// Criar a estrutura de dados para o empréstimo
		emprestimo := Emprestimo{
			Codigo:         codigo,
			CPF_Professor:  cpfProfessor,
			Nome_Professor: nomeProfessor,
			Horario_inicio: horarioInicio,
			Horario_fim:    horarioFim,
		}

		// Converter a estrutura de dados para JSON
		payload, err := json.Marshal(emprestimo)
		if err != nil {
			http.Error(w, "Erro ao converter os dados para JSON", http.StatusInternalServerError)
			log.Println("Erro ao converter os dados para JSON:", err) // Adiciona um log de erro
			return
		}

		// Enviar a solicitação HTTP POST para o backend
		resp, err := http.Post("http://localhost:8080/emprestimos", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			http.Error(w, "Erro ao enviar a solicitação HTTP para o backend", http.StatusInternalServerError)
			log.Println("Erro ao enviar a solicitação HTTP para o backend:", err) // Adiciona um log de erro
			return
		}
		defer resp.Body.Close()

		// Verificar a resposta do backend
		if resp.StatusCode == http.StatusCreated {
			fmt.Fprintf(w, "Empréstimo cadastrado com sucesso")
		} else {
			http.Error(w, "Erro ao cadastrar o empréstimo", http.StatusInternalServerError)
			log.Println("Erro ao cadastrar o empréstimo: status", resp.StatusCode) // Adiciona um log de erro
		}
	}
}
