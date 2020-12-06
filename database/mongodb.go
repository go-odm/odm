/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 9/4/20 5:13 PM
*/

package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	godm "github.com/go-odm/odm"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	collection *mongo.Collection
}

func handlePipeline(pipeline func(interface{}) error, input interface{}) error {
	if err := pipeline(input); err != nil {
		return err
	}
	return nil
}

func (m mongodb) InsertOne(ctx context.Context, doc interface{}, pipeline godm.PipelineFn, opts *options.InsertOneOptions) error {
	result, err := m.collection.InsertOne(ctx, doc, opts)
	if err != nil {
		return err
	}
	return handlePipeline(pipeline, result)
}

func (m mongodb) Find(ctx context.Context, filter interface{}, pipeline godm.PipelineFn, opts *options.FindOptions) error {
	result, err := m.collection.Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	return handlePipeline(pipeline, result)
}

func (m mongodb) FindOne(ctx context.Context, filter interface{}, pipeline godm.PipelineFn, opts *options.FindOneOptions) error {
	result := m.collection.FindOne(ctx, filter, opts)
	if result.Err() != nil {
		return result.Err()
	}
	return handlePipeline(pipeline, result)
}

func (m mongodb) FindOneAndDelete(ctx context.Context, filter interface{}, pipeline godm.PipelineFn, opts *options.FindOneAndDeleteOptions) error {
	result := m.collection.FindOneAndDelete(ctx, filter, opts)
	if result.Err() != nil {
		return result.Err()
	}
	return handlePipeline(pipeline, result)
}

func (m mongodb) FindOneAndUpdate(ctx context.Context, filter, update interface{}, pipeline godm.PipelineFn, opts *options.FindOneAndUpdateOptions) error {
	result := m.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if result.Err() != nil {
		return result.Err()
	}
	return handlePipeline(pipeline, result)
}

func (m mongodb) FindOneAndReplace(ctx context.Context, filter, replacement interface{}, pipeline godm.PipelineFn, opts *options.FindOneAndReplaceOptions) error {
	result := m.collection.FindOneAndReplace(ctx, filter, replacement, opts)
	if result.Err() != nil {
		return result.Err()
	}
	return handlePipeline(pipeline, result)
}

func (m mongodb) Where(args ...interface{}) *godm.Query {
	q := godm.NewQueryWitConn(m)
	q.Where(args)
	return q
}

func newConnection(c *mongo.Collection) godm.Connection {
	return mongodb{collection: c}
}
