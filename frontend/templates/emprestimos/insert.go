package emprestimos

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
