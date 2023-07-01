package emprestimos

import (
	"net/http"
	"log"
	"encoding/json"
	"html/template"
	"io/ioutil"
)

type EmprestimoData struct {
	Key   string      `json:"Key"`
	Value interface{} `json:"Value"`
}

type EmprestimoResponse [][]EmprestimoData

func Read(w http.ResponseWriter, r *http.Request) {
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
						<td><a class="button-update" href="/atualizar-emprestimo?codigo={{index . 1}}">Atualizar</a></td>
						<td><a class="button-delete" href="/deletar-emprestimos?codigo={{index . 1}}">Deletar</a></td>
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
