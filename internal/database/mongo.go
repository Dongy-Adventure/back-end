package database

import (
	"context"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDatabase(conf *config.DbConfig) (db *mongo.Database, err error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.MongoURL))
	if err != nil {
		return nil, err
	}
	database := client.Database("MarketPlace")
	return database, nil
}
