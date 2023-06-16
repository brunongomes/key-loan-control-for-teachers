package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Disciplina struct {
	Codigo       string `json:"codigo"`
	Nome         string `json:"nome"`
	CargaHoraria int    `json:"cargaHoraria"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// Exibir o formulário HTML para o usuário
			html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Cadastrar Disciplina</title>
			</head>
			<body>
				<h1>Cadastrar Disciplina</h1>
				<form action="/" method="POST">
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
		} else if r.Method == "POST" {
			// Processar o formulário e enviar a solicitação HTTP para o backend

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
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
