package main

import (
	"log"
	"time"
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name  string
	Email string
}

func ConnectToMongoDB() (*mongo.Client, error) {
	// Defina as opções de conexão
	clientOptions := options.Client().ApplyURI("mongodb://root:12345@172.16.56.45:27017")

	// Conecte ao servidor MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
			return nil, err
	}

	// Verifique a conexão
	err = client.Ping(context.Background(), nil)
	if err != nil {
			return nil, err
	}

	return client, nil
}

func main() {
	client, err := ConnectToMongoDB()
	if err != nil {
			fmt.Println("Erro ao conectar ao MongoDB:", err)
			return
	}

	// Faça algo com o cliente do MongoDB, como executar consultas ou operações de gravação
	// Obter uma referência para o banco de dados
	database := client.Database("meu_banco_de_dados")

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

	// Feche a conexão quando não for mais necessária
	err = client.Disconnect(context.Background())
	if err != nil {
			fmt.Println("Erro ao desconectar do MongoDB:", err)
			return
	}
}
