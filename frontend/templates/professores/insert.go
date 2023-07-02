package professores

import (
	"fmt"
	"net/http"
	"strings"
	"bytes"
	"log"
	"encoding/json"

	"../../models"
)

func Insert(w http.ResponseWriter, r *http.Request) {
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
		professor := models.Professor{
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