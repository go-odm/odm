/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 9/4/20 5:13 PM
*/

package database

import (
	"context"

	godm "github.com/go-odm/odm"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	client *mongo.Client
}

func (m mongodb) InsertOne() {
	panic("implement me")
}

func (m mongodb) InsertMany() {
	panic("implement me")
}

func (m mongodb) Find(ctx context.Context, filter interface{}, opts interface{}) interface{} {
	panic("implement me")
}

func (m mongodb) FindOne() {
	panic("implement me")
}

func (m mongodb) FindOneAndDelete() {
	panic("implement me")
}

func (m mongodb) FindOneAndUpdate() {
	panic("implement me")
}

func (m mongodb) ReplaceOne() {
	panic("implement me")
}

func (m mongodb) UpdateOne() {
	panic("implement me")
}

func (m mongodb) UpdateMany() {
	panic("implement me")
}

func (m mongodb) DeleteOne() {
	panic("implement me")
}

func (m mongodb) DeleteMany() {
	panic("implement me")
}

func (m mongodb) Transaction() {
	panic("implement me")
}

func newConnection(c *mongo.Client) godm.Connection {
	return mongodb{client: c}
}
