package mongodb

import (
	"fmt"
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI("mongodb://root:12345@172.16.56.45:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }
    err = client.Ping(context.Background(), nil)
    if err != nil {
        return nil, err
    }
    return client, nil
}

func DisconnectFromMongoDB(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
			fmt.Println("Erro ao desconectar do MongoDB:", err)
			return
	}
	fmt.Println("Desconectado do MongoDB com sucesso!")
}

