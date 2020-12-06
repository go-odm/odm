/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2020/12/06 12:54
*/

package godm

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	skip       int32
	projection M
	sort       M
}

func (o *Options) FindOneOptions() *options.FindOneOptions {
	return nil
}

func (o *Options) FindOneAndReplaceOptions() *options.FindOneAndReplaceOptions {
	return nil
}
func (o *Options) FindOneAndUpdateOptions() *options.FindOneAndUpdateOptions {
	return nil
}
func (o *Options) FindOneAndDeleteOptions() *options.FindOneAndDeleteOptions {
	return nil
}
func (o *Options) InsertOneOptions() *options.InsertOneOptions {
	return nil
}
