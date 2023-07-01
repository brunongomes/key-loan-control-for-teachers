
package disciplinas

import (
	"net/http"
	"log"
	"encoding/json"
	"html/template"
	"io/ioutil"
)

type DisciplinaData struct {
	Key   string `json:"Key"`
	Value interface{} `json:"Value"`
}

type DisciplinaResponse [][]DisciplinaData

func Read(w http.ResponseWriter, r *http.Request) {
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
