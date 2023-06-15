package main

import (
	"log"
	"time"
	"fmt"
	"context"
	"./pkg/mongodb"
)

type Person struct {
	Name  string
	Email string
}

func main() {
	client, err := mongodb.ConnectToMongoDB()
	if err != nil {
			fmt.Println("Erro ao conectar ao MongoDB:", err)
			return
	}

	// Faça algo com o cliente do MongoDB, como executar consultas ou operações de gravação
	// Obter uma referência para o banco de dados
	database := client.Database("key_loan")

	// Obter uma referência para a coleção
	collection := database.Collection("minha_colecao")

	// Criar um contexto com timeout para as operações
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Criar um documento para inserir na coleção
	person := Person{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// Inserir o documento na coleção
	_, err = collection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documento inserido com sucesso!")

	mongodb.DisconnectFromMongoDB(client)
}
