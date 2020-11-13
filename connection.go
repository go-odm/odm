/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 9/4/20 5:12 PM
*/

package godm

import (
	"context"
)

type Connection interface {
	InsertOne()
	InsertMany()
	Find(ctx context.Context, filter interface{}, opts interface{}) interface{}
	FindOne()
	FindOneAndDelete()
	FindOneAndUpdate()
	ReplaceOne()
	UpdateOne()
	UpdateMany()
	DeleteOne()
	DeleteMany()
	Transaction()
}
