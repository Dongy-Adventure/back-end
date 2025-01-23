package databases

import (
	"context"
	"log"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDatabase(conf *config.DbConfig) (db *mongo.Database, err error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.MongoURL))
	if err != nil {
		return nil, err
	}
	database := client.Database("MarketPlace")

	collection := database.Collection("testCollection")

	// Inserting a dummy document to ensure the collection is created
	_, err = collection.InsertOne(context.Background(), bson.D{{Key: "name", Value: "example"}})
	if err != nil {
		log.Println("Error inserting document:", err)
	}
	return database, nil
}
