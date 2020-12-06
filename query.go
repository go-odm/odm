/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 9/3/20 6:59 PM
*/

package godm

import (
	"context"
	"fmt"
	"strings"
)

func initQuery() *Query {
	q := Query{
		operation:  "",
		options:    nil,
		conditions: make(M),
		field:      "",
	}
	return &q
}

func NewQuery(fn ...func(query *Query)) *Query {
	q := Query{
		operation: "",
		options: &Options{
			skip:       0,
			projection: make(M),
		},
		conditions: make(M),
		field:      "",
	}
	if len(fn) > 0 {
		for _, f := range fn {
			f(&q)
		}
	}
	return &q
}

func SetConn(c Connection) func(query *Query) {
	return func(query *Query) {
		query.connection = c
	}
}

func NewQueryWitConn(c Connection) *Query {
	return NewQuery(SetConn(c))
}

type Query struct {
	connection Connection
	operation  string
	options    *Options
	conditions M
	field      string
	model      interface{}
}

//TODO:
//func (q *Query) Find(ctx context.Context, model interface{}) error {
//	//res := q.connection.Find(ctx, q.conditions, q.options)
//	//err := json.Unmarshal(res, model)
//	//return err
//}

func (q *Query) FindOne(ctx context.Context, pipeline PipelineFn) error {
	return q.connection.FindOne(ctx, q.conditions, pipeline, q.options.FindOneOptions())
}
func (q *Query) InsertOne(ctx context.Context, pipeline PipelineFn) error {
	return q.connection.InsertOne(ctx, q.conditions, pipeline, q.options.InsertOneOptions())
}

func (q *Query) FindOneAndDelete(ctx context.Context, pipeline PipelineFn) error {
	return q.connection.FindOneAndDelete(ctx, q.conditions, pipeline, q.options.FindOneAndDeleteOptions())
}

func (q *Query) FindOneAndUpdateOne(ctx context.Context, update interface{}, pipeline PipelineFn) error {
	return q.connection.FindOneAndUpdate(ctx, q.conditions, update, pipeline, q.options.FindOneAndUpdateOptions())
}

func (q *Query) FindOneAndReplace(ctx context.Context, replacement interface{}, pipeline PipelineFn) error {
	return q.connection.FindOneAndReplace(ctx, q.conditions, replacement, pipeline, q.options.FindOneAndReplaceOptions())
}

func (q *Query) GetConditions() M {
	return q.conditions
}

func (q *Query) GetSelect() M {
	return q.options.projection
}

func mappingStringToFieldSets(value Input, projection bool) Input {
	out := -1
	if projection {
		out = 0
	}
	obj := make(M)
	switch value.(type) {
	case string:
		strArray := strings.Fields(strings.TrimSpace(value.(string)))
		for _, v := range strArray {
			if v[0] == '-' {
				v = v[1:]
				obj[v] = out
			} else {
				obj[v] = 1
			}
		}
	case M:
		obj = value.(M)
	}
	return obj
}

func (q *Query) Sort(value interface{}) *Query {
	q.options.sort = mappingStringToFieldSets(value, false).(M)
	return q
}

func (q *Query) Select(value interface{}) *Query {
	q.options.projection = mappingStringToFieldSets(value, true).(M)
	return q
}

func (q *Query) Skip(value int32) *Query {
	q.options.skip = value
	return q
}

func (q *Query) Where(args ...interface{}) *Query {
	//q.field = field
	switch len(args) {
	// first args is string
	case 1:
		if field, ok := args[0].(string); ok {
			q.Set(field)
		}
		// Where("version",1) where version is equals q
	case 2:
		if field, ok := args[0].(string); ok {
			q.Set(field).Eq(args[1])
		}
		// Where("version",">=",1) gte 1
	case 3:
		if field, ok := args[0].(string); ok {
			q.Set(field)
		}
		if operators, ok := args[1].(string); ok {
			q.bindCondition(
				chain(
					getFlagWrapperFromString(operators),
					inputBuilder,
				)(args[2]),
			)
		}
	}
	return q
}

func (q *Query) Set(field string) *Query {
	q.field = field
	return q
}

func (q *Query) Eq(input interface{}) *Query {
	q.
		ensureField("Eq").
		bindCondition(
			chain(
				inputLogger,
				inputBuilder,
			)(input),
		).
		resetField()
	return q
}

func (q *Query) Equals(input interface{}) *Query {
	q.
		ensureField("Equals").
		bindCondition(
			chain(inputBuilder)(input),
		).
		resetField()
	return q
}

func (q *Query) AutoBindConditions(flag string, condition interface{}) *Query {
	if q.hasField() {
		q.bindCondition(
			chain(
				inputBuilder,
			)(condition),
		).resetField()
	} else {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(condition),
		).resetField()
	}
	return q
}

/*
	e.g. query.Or([]M{{"name": "weny"}, {"age": "20"}})
*/
func (q *Query) Or(value interface{}) *Query {
	flag := "$or"
	return q.AutoBindConditions(flag, value)
}

/*
	e.g. query.Nor([]M{{"name": "weny"}, {"age": "20"}})
*/
func (q *Query) Nor(value interface{}) *Query {
	flag := "$nor"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) And(value interface{}) *Query {
	flag := "$and"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Not(value interface{}) *Query {
	flag := "$not"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Gt(value interface{}) *Query {
	flag := "$gt"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Gte(value interface{}) *Query {
	flag := "$gte"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Lt(value interface{}) *Query {
	flag := "$lt"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Lte(value interface{}) *Query {
	flag := "$lte"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Ne(value interface{}) *Query {
	flag := "$ne"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) In(value interface{}) *Query {
	flag := "$in"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Nin(value interface{}) *Query {
	flag := "$nin"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) All(value interface{}) *Query {
	flag := "$all"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Regex(value interface{}) *Query {
	flag := "$regex"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Size(value interface{}) *Query {
	flag := "$size"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) MaxDistance(value interface{}) *Query {
	flag := "$maxDistance"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) MinDistance(value interface{}) *Query {
	flag := "$minDistance"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) Mod(value interface{}) *Query {
	flag := "$mod"
	return q.AutoBindConditions(flag, value)
}

//TODO: geometry

func (q *Query) Exists(value bool) *Query {
	flag := "$exists"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) ElemMatch(value interface{}) *Query {
	flag := "$elemMatch"
	return q.AutoBindConditions(flag, value)
}

func (q *Query) bindCondition(value Input) *Query {
	q.conditions[q.field] = value
	return q
}

func (q *Query) resetField() *Query {
	q.field = ""
	return q
}

func (q *Query) setField(field string) *Query {
	q.field = field
	return q
}
func (q *Query) hasField() bool {
	if q.field == "" {
		return false
	}
	return true
}

func (q *Query) ensureField(method string) *Query {
	if q.field == "" {
		panic(method + " must be used after Where() ")
	}
	return q
}

func inputLogger(input Input) Input {
	go func() {
		fmt.Print(input)
	}()
	return input
}

func inputWrapper(flag string) InputEndpoint {
	return func(input Input) Input {
		return M{flag: input}
	}
}
func getFlagWrapperFromString(arg string) InputEndpoint {
	switch arg {
	case "<":
		return inputWrapper("$lt")
	case ">":
		return inputWrapper("$gt")
	case "=":
		return inputWrapper("$eq")
	case ">=":
		return inputWrapper("$gte")
	case "<=":
		return inputWrapper("$lte")
	default:
		return inputWrapper("$eq")
	}
}

func inputBuilder(input Input) Input {
	var res interface{}
	switch input.(type) {
	case func(q *Query):
		query := NewQuery()
		input.(func(q *Query))(query)
		res = query.conditions
		break
	case func(q *Query) *Query:
		res = input.(func(q *Query) *Query)(initQuery()).conditions
		break
	case interface{}:
		res = input
		break
	}
	return res
}
