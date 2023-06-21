package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"bytes"
	"log"
	"encoding/json"
	"html/template"
	"io/ioutil"
)
type Emprestimo struct {
	Codigo         int    `json:"codigo"`
	CPF_Professor  string `json:"cpfProfessor"`
	Nome_Professor string `json:"nomeProfessor"`
	Horario_inicio string `json:"horarioInicio"`
	Horario_fim    string `json:"horarioFim"`
}

type EmprestimoData struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
}

type EmprestimoResponse [][]EmprestimoData

func EmprestimosHandler(w http.ResponseWriter, r *http.Request) {
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

func CadastrarEmprestimoHandler(w http.ResponseWriter, r *http.Request) {
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
			mensagem := "Empréstimo cadastrado com sucesso"
			html := fmt.Sprintf(`<script>alert("%s"); window.location.href = "/emprestimos";</script>`, mensagem)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, html)
		} else {
			mensagem := "Erro ao cadastrar a empréstimo"
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, `<script>alert("%s");</script>`, mensagem)
			log.Println("Erro ao cadastrar a empréstimo: status", resp.StatusCode) // Adiciona um log de erro
		}
	}
}

func VisualizarEmprestimosHandler(w http.ResponseWriter, r *http.Request) {
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
						<th>Atualizar</th>
						<th>Deletar</th>
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
						<td><a href="/atualizar-disciplina?codigo={{index . 1}}">Atualizar</a></td>
						<td><a href="/disciplinas?codigo={{index . 1}}">Deletar</a></td>
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
