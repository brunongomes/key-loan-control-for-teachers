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
	"io/ioutil"

	"./handlers"
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

	router.HandleFunc("/", handlers.Home).Methods("GET")
	router.HandleFunc("/disciplinas", disciplinasHandler).Methods("GET")
	router.HandleFunc("/visualizar-disciplinas", visualizarDisciplinasHandler).Methods("GET") // Nova rota para visualizar disciplinas 
	router.HandleFunc("/visualizar-professores", visualizarProfessoresHandler).Methods("GET") // Nova rota para visualizar professores 
	router.HandleFunc("/visualizar-emprestimos", visualizarEmprestimosHandler).Methods("GET") // Nova rota para visualizar emprestimos 
	router.HandleFunc("/emprestimos", emprestimosHandler).Methods("GET")
	router.HandleFunc("/professores", professoresHandler).Methods("GET")
	router.HandleFunc("/disciplinas", cadastrarDisciplinaHandler).Methods("POST")
	router.HandleFunc("/professores", cadastrarProfessorHandler).Methods("POST")
	router.HandleFunc("/emprestimos", cadastrarEmprestimoHandler).Methods("POST")

	router.HandleFunc("/deletar-disciplina/{codigo}", deletarDisciplinaHandler).Methods("DELETE")

	
	log.Fatal(http.ListenAndServe(":8081", router))
}

func disciplinasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Exibir a página de disciplinas
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Cadastrar Disciplina</title>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body class="container">
			<h1>Cadastrar Disciplina</h1>
			<div class="form-container">
				<form action="/disciplinas" method="POST">
					<label for="codigo">Código:</label><br>
					<input type="text" id="codigo" name="codigo"><br><br>
					<label for="nome">Nome:</label><br>
					<input type="text" id="nome" name="nome"><br><br>
					<label for="cargaHoraria">Carga Horária:</label><br>
					<input type="text" id="cargaHoraria" name="cargaHoraria"><br><br>
					<input class="button-form" type="submit" value="Cadastrar">
				</form>
			</div>
			<br>
			<form action="/" method="GET">
				<input class="button-form" type="submit" value="Voltar">
			</form>
			<br>
			<form action="/visualizar-disciplinas" method="GET">
				<input class="button-form" type="submit" value="Visualizar">
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
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body class="container" >
			<h1>Cadastrar Empréstimos</h1>
			<div class="form-container">
				<form action="/emprestimos" method="POST">
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
					<input class="button-form" type="submit" value="Cadastrar">
				</form>
				</div>
				<br>
				<form action="/" method="GET">
					<input class="button-form" type="submit" value="Voltar">
				</form>
				<br>
				<form action="/visualizar-emprestimos" method="GET">
					<input class="button-form" type="submit" value="Visualizar">
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
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body class="container" >
			<h1>Cadastrar Professor</h1>
			<div class="form-container">
				<form action="/professores" method="POST">
					<label for="cpf">CPF:</label><br>
					<input type="text" id="cpf" name="cpf"><br><br>
					<label for="nome">Nome:</label><br>
					<input type="text" id="nome" name="nome"><br><br>
					<input class="button-form" type="submit" value="Cadastrar">
				</form>
				</div>
				<br>
				<form action="/" method="GET">
					<input class="button-form" type="submit" value="Voltar">
				</form>
				<br>
				<form action="/visualizar-professores" method="GET">
					<input class="button-form" type="submit" value="Visualizar">
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
				mensagem := "Disciplina cadastrada com sucesso"
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				fmt.Fprintf(w, `<script>alert("%s");</script>`, mensagem)
			} else {
				mensagem := "Erro ao cadastrar a disciplina"
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				fmt.Fprintf(w, `<script>alert("%s");</script>`, mensagem)
				log.Println("Erro ao cadastrar a disciplina: status", resp.StatusCode) // Adiciona um log de erro
			}
	}
}

type DisciplinaData struct {
	Key   string `json:"Key"`
	Value interface{} `json:"Value"`
}

type DisciplinaResponse [][]DisciplinaData

func visualizarDisciplinasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Fazer a solicitação HTTP GET para obter os dados das disciplinas do backend
		resp, err := http.Get("http://localhost:8080/disciplinas")
		if err != nil {
			http.Error(w, "Erro ao obter os dados das disciplinas", http.StatusInternalServerError)
			log.Println("Erro ao obter os dados das disciplinas:", err)
			return
		}
		defer resp.Body.Close()

		// Ler a resposta do backend
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Erro ao ler a resposta do backend", http.StatusInternalServerError)
			log.Println("Erro ao ler a resposta do backend:", err)
			return
		}

		// Processar os dados da resposta em formato JSON
		var disciplinas DisciplinaResponse
		err = json.Unmarshal(body, &disciplinas)
		if err != nil {
			http.Error(w, "Erro ao processar os dados das disciplinas", http.StatusInternalServerError)
			log.Println("Erro ao processar os dados das disciplinas:", err)
			return
		}

		// Exibir a página de visualização das disciplinas
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Visualizar Disciplinas</title>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body>
			<h1>Visualizar Disciplinas</h1>
			<table class="form-tabela">
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
						<td>{{index . 1}}</td>
						<td>{{index . 2}}</td>
						<td>{{index . 3}}</td>
						<td><a href="/atualizar-disciplina?codigo={{index . 1}}">Atualizar</a></td>
						<td><a href="/deletar-disciplina?codigo={{index . 1}}">Deletar</a></td>
					</tr>
					{{end}}
				</tbody>
			</table>
			<br>
			<form action="/disciplinas" method="GET">
				<input class="button-form" type="submit" value="Voltar">
			</form>
		</body>
		</html>
		`

		// Renderizar o HTML substituindo os dados das disciplinas
		tmpl, err := template.New("visualizar_disciplinas").Parse(html)
		if err != nil {
			http.Error(w, "Erro ao renderizar o HTML", http.StatusInternalServerError)
			log.Println("Erro ao renderizar o HTML:", err)
			return
		}

		err = tmpl.Execute(w, disciplinas)
		if err != nil {
			http.Error(w, "Erro ao renderizar o HTML", http.StatusInternalServerError)
			log.Println("Erro ao renderizar o HTML:", err)
			return
		}
	}
}

func deletarDisciplinaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		vars := mux.Vars(r)
		codigo := vars["codigo"]

		// Montar a URL para a solicitação HTTP DELETE
		url := "http://localhost:8080/disciplinas/" + codigo

		// Criar a solicitação HTTP DELETE
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			http.Error(w, "Erro ao criar a solicitação HTTP", http.StatusInternalServerError)
			log.Println("Erro ao criar a solicitação HTTP:", err)
			return
		}

		// Enviar a solicitação HTTP DELETE para o backend
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Erro ao enviar a solicitação HTTP para o backend", http.StatusInternalServerError)
			log.Println("Erro ao enviar a solicitação HTTP para o backend:", err)
			return
		}
		defer resp.Body.Close()

		// Verificar a resposta do backend
		if resp.StatusCode == http.StatusOK {
			// Redirecionar de volta para a página de visualização das disciplinas
			http.Redirect(w, r, "/visualizar-disciplinas", http.StatusSeeOther)
		} else {
			http.Error(w, "Erro ao deletar a disciplina", http.StatusInternalServerError)
			log.Println("Erro ao deletar a disciplina: status", resp.StatusCode)
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


type ProfessorData struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
}

type ProfessorResponse [][]ProfessorData

func visualizarProfessoresHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Fazer a solicitação HTTP GET para obter os dados dos professores do backend
		resp, err := http.Get("http://localhost:8080/professores")
		if err != nil {
			http.Error(w, "Erro ao obter os dados dos professores", http.StatusInternalServerError)
			log.Println("Erro ao obter os dados dos professores:", err)
			return
		}
		defer resp.Body.Close()

		// Ler a resposta do backend
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Erro ao ler a resposta do backend", http.StatusInternalServerError)
			log.Println("Erro ao ler a resposta do backend:", err)
			return
		}

		// Processar os dados da resposta em formato JSON
		var professores ProfessorResponse
		err = json.Unmarshal(body, &professores)
		if err != nil {
			http.Error(w, "Erro ao processar os dados dos professores", http.StatusInternalServerError)
			log.Println("Erro ao processar os dados dos professores:", err)
			return
		}

		// Exibir a página de visualização dos professores
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Visualizar Professores</title>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body>
			<h1>Visualizar Professores</h1>
			<table class="form-tabela">
				<thead>
					<tr>
						<th>CPF</th>
						<th>Nome</th>
					</tr>
				</thead>
				<tbody>
					{{range .}}
					<tr>
						<td>{{index . 1}}</td>
						<td>{{index . 2}}</td>
					</tr>
					{{end}}
				</tbody>
			</table>
			<br>
			<form action="/professores" method="GET">
				<input class="button-form" type="submit" value="Voltar">
			</form>
		</body>
		</html>
		`

		// Renderizar o HTML substituindo os dados dos professores
		tmpl, err := template.New("visualizar_professores").Parse(html)
		if err != nil {
			http.Error(w, "Erro ao renderizar o HTML", http.StatusInternalServerError)
			log.Println("Erro ao renderizar o HTML:", err)
			return
		}

		err = tmpl.Execute(w, professores)
		if err != nil {
			http.Error(w, "Erro ao renderizar o HTML", http.StatusInternalServerError)
			log.Println("Erro ao renderizar o HTML:", err)
			return
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


type EmprestimoData struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
}

type EmprestimoResponse [][]EmprestimoData


func visualizarEmprestimosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Fazer a solicitação HTTP GET para obter os dados dos empréstimos do backend
		resp, err := http.Get("http://localhost:8080/emprestimos")
		if err != nil {
			http.Error(w, "Erro ao obter os dados dos empréstimos", http.StatusInternalServerError)
			log.Println("Erro ao obter os dados dos empréstimos:", err)
			return
		}
		defer resp.Body.Close()

		// Ler a resposta do backend
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Erro ao ler a resposta do backend", http.StatusInternalServerError)
			log.Println("Erro ao ler a resposta do backend:", err)
			return
		}

		// Processar os dados da resposta em formato JSON
		var emprestimos [][]interface{}
		err = json.Unmarshal(body, &emprestimos)
		if err != nil {
			http.Error(w, "Erro ao processar os dados dos empréstimos", http.StatusInternalServerError)
			log.Println("Erro ao processar os dados dos empréstimos:", err)
			return
		}

		// Exibir a página de visualização dos empréstimos
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Visualizar Empréstimos</title>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body>
			<h1>Visualizar Empréstimos</h1>
			<table class="form-tabela">
				<thead>
					<tr>
						<th>Código</th>
						<th>CPF Professor</th>
						<th>Nome Professor</th>
						<th>Horário Início</th>
						<th>Horário Fim</th>
					</tr>
				</thead>
				<tbody>
					{{range .}}
					<tr>
						<td>{{index . 0}}</td>
						<td>{{index . 1}}</td>
						<td>{{index . 2}}</td>
						<td>{{index . 3}}</td>
						<td>{{index . 4}}</td>
					</tr>
					{{end}}
				</tbody>
			</table>
			<br>
			<form action="/emprestimos" method="GET">
				<input class="button-form" type="submit" value="Voltar">
			</form>
		</body>
		</html>
		`

		// Renderizar o HTML substituindo os dados dos empréstimos
		tmpl, err := template.New("visualizar_emprestimos").Parse(html)
		if err != nil {
			http.Error(w, "Erro ao renderizar o HTML", http.StatusInternalServerError)
			log.Println("Erro ao renderizar o HTML:", err)
			return
		}

		err = tmpl.Execute(w, emprestimos)
		if err != nil {
			http.Error(w, "Erro ao renderizar o HTML", http.StatusInternalServerError)
			log.Println("Erro ao renderizar o HTML:", err)
			return
		}
	}
}
