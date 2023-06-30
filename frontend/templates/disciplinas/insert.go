package disciplinas

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"bytes"
	"log"
	"encoding/json"
)

func Insert(w http.ResponseWriter, r *http.Request) {
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
