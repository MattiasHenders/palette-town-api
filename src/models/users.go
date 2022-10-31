package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName *string            `json:"firstName,omitempty" bson:"firstName"`
	LastName  *string            `json:"lastName,omitempty" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	Password  *string            `json:"-" bson:"password"`
	UserType  string             `json:"userType" bson:"userType"`
	Credits   int                `json:"credits" bson:"credits"`
}
