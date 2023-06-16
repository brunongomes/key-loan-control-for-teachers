package disciplinas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
	"../../pkg/mongodb"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func CadastrarDisciplina(w http.ResponseWriter, r *http.Request) {
	// Decodifica o JSON recebido na requisição em uma struct Disciplina
	var disciplina Disciplina
	err := json.NewDecoder(r.Body).Decode(&disciplina)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}

	// Insere a disciplina no banco de dados
	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	err = db.Insert("disciplinas", disciplina)
	if err != nil {
		http.Error(w, "Erro ao cadastrar a disciplina", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Disciplina cadastrada com sucesso")
}

// Handler para listar disciplinas
func ListarDisciplinas(w http.ResponseWriter, r *http.Request) {
	// Lógica de listagem de disciplinas
	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	disciplinas, err := db.Read("disciplinas", bson.M{})
	if err != nil {
		http.Error(w, "Erro ao listar as disciplinas", http.StatusInternalServerError)
		return
	}

	// Serializa as disciplinas em JSON
	jsonBytes, err := json.Marshal(disciplinas)
	if err != nil {
		http.Error(w, "Erro ao serializar as disciplinas em JSON", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho de resposta como JSON
	w.Header().Set("Content-Type", "application/json")
	// Escreve os dados no corpo da resposta
	w.Write(jsonBytes)
}


// Handler para excluir uma disciplina
func ExcluirDisciplina(w http.ResponseWriter, r *http.Request) {
	// Lógica de exclusão de disciplina
	params := mux.Vars(r)
	codigo := params["codigo"]

	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao MongoDB", http.StatusInternalServerError)
		return
	}
	defer db.Client.Disconnect(context.Background())

	err = db.Delete("disciplinas", bson.M{"codigo": codigo})
	if err != nil {
		http.Error(w, "Erro ao excluir a disciplina", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Disciplina excluída com sucesso")
}

// Handler para atualizar uma disciplina
func AtualizarDisciplina(w http.ResponseWriter, r *http.Request) {
	// Lógica de atualização de disciplina
	params := mux.Vars(r)
	codigo := params["codigo"]

	// Decodifica o JSON recebido na requisição em uma struct Disciplina
	var disciplina Disciplina
	err := json.NewDecoder(r.Body).Decode(&disciplina)
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
			"nome":         disciplina.Nome,
			"cargaHoraria": disciplina.CargaHoraria,
		},
	}

	err = db.Update("disciplinas", bson.M{"codigo": codigo}, updateFields)
	if err != nil {
		http.Error(w, "Erro ao atualizar a disciplina", http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Disciplina atualizada com sucesso")
}
