package emprestimos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
	"../../pkg/mongodb"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func CadastrarEmprestimo(w http.ResponseWriter, r *http.Request) {
	// Decodifica o JSON recebido na requisição em uma struct Emprestimo
	var emprestimo Emprestimo
	err := json.NewDecoder(r.Body).Decode(&emprestimo)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}

	// Insere o empréstimo no banco de dados
	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	err = db.Insert("emprestimos", emprestimo)
	if err != nil {
		http.Error(w, "Erro ao cadastrar o empréstimo", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Empréstimo cadastrado com sucesso")
}

// Handler para listar empréstimos
func ListarEmprestimos(w http.ResponseWriter, r *http.Request) {
	// Lógica de listagem de empréstimos
	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	emprestimos, err := db.Read("emprestimos", bson.M{})
	if err != nil {
		http.Error(w, "Erro ao listar os empréstimos", http.StatusInternalServerError)
		return
	}

	// Serializa os empréstimos em JSON
	jsonBytes, err := json.Marshal(emprestimos)
	if err != nil {
		http.Error(w, "Erro ao serializar os empréstimos em JSON", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho de resposta como JSON
	w.Header().Set("Content-Type", "application/json")
	// Escreve os dados no corpo da resposta
	w.Write(jsonBytes)
}

// Handler para excluir um empréstimo
func ExcluirEmprestimo(w http.ResponseWriter, r *http.Request) {
	// Lógica de exclusão de empréstimo
	params := mux.Vars(r)
	codigo := params["codigo"]

	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	err = db.Delete("emprestimos", bson.M{"codigo": codigo})
	if err != nil {
		http.Error(w, "Erro ao excluir o empréstimo", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Empréstimo excluído com sucesso")
}

// Handler para atualizar um empréstimo
func AtualizarEmprestimo(w http.ResponseWriter, r *http.Request) {
	// Lógica de atualização de empréstimo
	params := mux.Vars(r)
	codigo := params["codigo"]

	// Decodifica o JSON recebido na requisição em uma struct Emprestimo
	var emprestimo Emprestimo
	err := json.NewDecoder(r.Body).Decode(&emprestimo)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}

	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	updateFields := bson.M{
		"$set": bson.M{
			"cpfProfessor":  emprestimo.CPF_Professor,
			"nomeProfessor": emprestimo.Nome_Professor,
			"horarioInicio": emprestimo.Horario_inicio,
			"horarioFim":    emprestimo.Horario_fim,
		},
	}

	err = db.Update("emprestimos", bson.M{"codigo": codigo}, updateFields)
	if err != nil {
		http.Error(w, "Erro ao atualizar o empréstimo", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Empréstimo atualizado com sucesso")
}
