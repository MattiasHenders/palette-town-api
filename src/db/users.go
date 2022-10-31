package db

import (
	"net/http"

	"github.com/MattiasHenders/palette-town-api/src/internal/errors"
	"github.com/MattiasHenders/palette-town-api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollection = "users"
)

func GetUserByID(id string) (*models.User, *errors.HTTPError) {

	collection := Mongo.Client.Database(Mongo.DBName).Collection(userCollection)

	var result models.User
	err := collection.FindOne(Mongo.Context, bson.D{{Key: "_id", Value: id}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewHTTPError(err, http.StatusNotFound, err.Error())
		}
		return nil, errors.NewHTTPError(err, http.StatusInternalServerError, err.Error())
	}

	return &result, nil
}

func GetUserByEmail(email string) (*models.User, *errors.HTTPError) {

	collection := Mongo.Client.Database(Mongo.DBName).Collection(userCollection)

	var result models.User
	err := collection.FindOne(Mongo.Context, bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewHTTPError(err, http.StatusNotFound, err.Error())
		}
		return nil, errors.NewHTTPError(err, http.StatusInternalServerError, err.Error())
	}

	return &result, nil
}

func CreateNewUser(doc models.User) (*models.User, *errors.HTTPError) {

	collection := Mongo.Client.Database(Mongo.DBName).Collection(userCollection)

	_, err := collection.InsertOne(Mongo.Context, doc)
	if err != nil {
		return nil, errors.NewHTTPError(err, http.StatusInternalServerError, err.Error())
	}

	return &doc, nil
}
