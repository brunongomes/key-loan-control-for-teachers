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
	"github.com/gorilla/mux"
)

type Disciplina struct {
	Codigo       string `json:"codigo"`
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"cargaHoraria"`
}

type DisciplinaData struct {
	Key   string `json:"Key"`
	Value interface{} `json:"Value"`
}

type DisciplinaResponse [][]DisciplinaData

func DisciplinasHandler(w http.ResponseWriter, r *http.Request) {
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

func CadastrarDisciplinaHandler(w http.ResponseWriter, r *http.Request) {
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
				html := fmt.Sprintf(`<script>alert("%s"); window.location.href = "/disciplinas";</script>`, mensagem)
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				fmt.Fprintf(w, html)
			} else {
				mensagem := "Erro ao cadastrar a disciplina"
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				fmt.Fprintf(w, `<script>alert("%s");</script>`, mensagem)
				log.Println("Erro ao cadastrar a disciplina: status", resp.StatusCode) // Adiciona um log de erro
			}
	}
}

func VisualizarDisciplinasHandler(w http.ResponseWriter, r *http.Request) {
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
						<td><a class="button-update" href="/atualizar-disciplina?codigo={{index . 1}}">Atualizar</a></td>
						<td><a class="button-delete" href="/deletar-disciplina?codigo={{index . 1}}">Deletar</a></td>
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

func DeletarDisciplinaHandler(w http.ResponseWriter, r *http.Request) {
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
