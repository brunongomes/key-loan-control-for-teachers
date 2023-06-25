package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"bytes"
	"log"
	"encoding/json"
	"html/template"
	"io/ioutil"
)

type Professor struct {
	CPF  string `json:"cpf"`
	Nome string `json:"nome"`
}

type ProfessorData struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
}

type ProfessorResponse [][]ProfessorData

func ProfessoresHandler(w http.ResponseWriter, r *http.Request) {
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
				<div class="btn"> 
					<a href="/visualizar-professores" class="button-form" method="GET" >Visualizar</a>
					<a href="/" class="button-form" method="GET" >Voltar</a>
				</div>
		</body>
		</html>
			`
		fmt.Fprintf(w, html)
	}
}

func CadastrarProfessorHandler(w http.ResponseWriter, r *http.Request) {
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
			mensagem := "Professor cadastrado com sucesso"
			html := fmt.Sprintf(`<script>alert("%s"); window.location.href = "/professores";</script>`, mensagem)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, html)
		} else {
			mensagem := "Erro ao cadastrar a professor"
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, `<script>alert("%s");</script>`, mensagem)
			log.Println("Erro ao cadastrar a professor: status", resp.StatusCode) // Adiciona um log de erro
		}
	}
}

func VisualizarProfessoresHandler(w http.ResponseWriter, r *http.Request) {
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
						<th>Atualizar</th>
						<th>Deletar</th>
					</tr>
				</thead>
				<tbody>
					{{range .}}
					<tr>
						<td>{{index . 1}}</td>
						<td>{{index . 2}}</td>
						<td><a class="button-update" href="/atualizar-professor?codigo={{index . 1}}">Atualizar</a></td>
						<td><a class="button-delete" href="/deletar-professores?codigo={{index . 1}}">Deletar</a></td>
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
