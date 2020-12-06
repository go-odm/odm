/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 9/4/20 5:12 PM
*/

package godm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection interface {
	InsertOne(ctx context.Context, doc interface{}, pipeline PipelineFn, opts *options.InsertOneOptions) error
	Find(ctx context.Context, filter interface{}, pipeline PipelineFn, opts *options.FindOptions) error
	FindOne(ctx context.Context, filter interface{}, pipeline PipelineFn, opts *options.FindOneOptions) error
	FindOneAndDelete(ctx context.Context, filter interface{}, pipeline PipelineFn, opts *options.FindOneAndDeleteOptions) error
	FindOneAndUpdate(ctx context.Context, filter, update interface{}, pipeline PipelineFn, opts *options.FindOneAndUpdateOptions) error
	FindOneAndReplace(ctx context.Context, filter, replacement interface{}, pipeline PipelineFn, opts *options.FindOneAndReplaceOptions) error
	Where(args ...interface{}) *Query
}
