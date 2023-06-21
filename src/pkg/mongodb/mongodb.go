package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type MongoDB struct {
	Client *mongo.Client
}

func ConnectToMongoDB() (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI("mongodb://root:12345@localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &MongoDB{Client: client}, nil
}

func (db *MongoDB) Insert(collectionName string, data interface{}) error {
	collection := db.Client.Database("meu_banco_de_dados").Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) Read(collectionName string, filter bson.M) ([]interface{}, error) {
	collection := db.Client.Database("meu_banco_de_dados").Collection(collectionName)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []interface{}
	for cursor.Next(context.Background()) {
		value := reflect.New(reflect.TypeOf(filter).Elem()).Interface()
		if err := cursor.Decode(value); err != nil {
			return nil, err
		}
		results = append(results, value)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (db *MongoDB) Update(collectionName string, filter bson.M, update bson.M) error {
	collection := db.Client.Database("meu_banco_de_dados").Collection(collectionName)
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) Delete(collectionName string, filter bson.M) error {
	collection := db.Client.Database("meu_banco_de_dados").Collection(collectionName)
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func DisconnectFromMongoDB(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
			fmt.Println("Erro ao desconectar do MongoDB:", err)
			return
	}
	fmt.Println("Desconectado do MongoDB com sucesso!")
}

