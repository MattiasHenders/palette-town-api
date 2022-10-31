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

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
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

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx := context.Background()

	// mongo.Connect return mongo.Client method
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(completeURI))

	mongoTemp := models.MongoDB{
		Client:  client,
		Context: ctx,
		DBName:  dbConfig.DBName,
	}

	Mongo = &mongoTemp
}

func Ping() error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := Mongo.Client.Ping(Mongo.Context, readpref.Primary()); err != nil {
		return err
	}
	return nil
}
