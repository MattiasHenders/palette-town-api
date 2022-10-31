package db

import (
	"context"
	"strings"

	"github.com/MattiasHenders/palette-town-api/config"
	"github.com/MattiasHenders/palette-town-api/src/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Mongo *models.MongoDB

func Close() {

	defer func() {
		if err := Mongo.Client.Disconnect(Mongo.Context); err != nil {
			panic(err)
		}
	}()
}

func Connect() {

	//Get DB config
	dbConfig := config.GetConfig().DB
	completeURI := strings.ReplaceAll(dbConfig.URI, "<USERNAME>", dbConfig.Username)
	completeURI = strings.ReplaceAll(completeURI, "<PASSWORD>", dbConfig.Password)

	ctx := context.Background()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(completeURI))

	mongoTemp := models.MongoDB{
		Client:  client,
		Context: ctx,
		DBName:  dbConfig.DBName,
	}

	Mongo = &mongoTemp
}

func Ping() error {

	if err := Mongo.Client.Ping(Mongo.Context, readpref.Primary()); err != nil {
		return err
	}
	return nil
}
