package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ServerResponse struct {
	Message     string `json:"message"`
	GivenInput  any    `json:"givenInput,omitempty"`
	Code        int    `json:"code"`
	Data        any    `json:"data"`
	CoolorsLink string `json:"coolorsLink"`
}

type ColorMindsResponse struct {
	Result [][]int `json:"result"`
}

type MongoDB struct {
	Client  *mongo.Client
	Context context.Context
	DBName  string
}
