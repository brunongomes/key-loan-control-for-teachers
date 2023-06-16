package professores

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
	"../../pkg/mongodb"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func CadastrarProfessor(w http.ResponseWriter, r *http.Request) {
	// Decodifica o JSON recebido na requisição em uma struct Professor
	var professor Professor
	err := json.NewDecoder(r.Body).Decode(&professor)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}

	// Insere o professor no banco de dados
	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	err = db.Insert("professores", professor)
	if err != nil {
		http.Error(w, "Erro ao cadastrar o professor", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Professor cadastrado com sucesso")
}

// Handler para listar professores
func ListarProfessores(w http.ResponseWriter, r *http.Request) {
	// Lógica de listagem de professores
	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	professores, err := db.Read("professores", bson.M{})
	if err != nil {
		http.Error(w, "Erro ao listar os professores", http.StatusInternalServerError)
		return
	}

	// Serializa os professores em JSON
	jsonBytes, err := json.Marshal(professores)
	if err != nil {
		http.Error(w, "Erro ao serializar os professores em JSON", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho de resposta como JSON
	w.Header().Set("Content-Type", "application/json")
	// Escreve os dados no corpo da resposta
	w.Write(jsonBytes)
}

// Handler para excluir um professor
func ExcluirProfessor(w http.ResponseWriter, r *http.Request) {
	// Lógica de exclusão de professor
	params := mux.Vars(r)
	cpf := params["cpf"]

	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	err = db.Delete("professores", bson.M{"cpf": cpf})
	if err != nil {
		http.Error(w, "Erro ao excluir o professor", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Professor excluído com sucesso")
}

// Handler para atualizar um professor
func AtualizarProfessor(w http.ResponseWriter, r *http.Request) {
	// Lógica de atualização de professor
	params := mux.Vars(r)
	cpf := params["cpf"]

	// Decodifica o JSON recebido na requisição em uma struct Professor
	var professor Professor
	err := json.NewDecoder(r.Body).Decode(&professor)
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
			"nome": professor.Nome,
		},
	}

	err = db.Update("professores", bson.M{"cpf": cpf}, updateFields)
	if err != nil {
		http.Error(w, "Erro ao atualizar o professor", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Professor atualizado com sucesso")
}
