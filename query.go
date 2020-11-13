/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 9/3/20 6:59 PM
*/

package godm

import (
	"fmt"
	"strings"
)

type Options struct {
	skip       int32
	projection M
	sort       M
}

func initQuery() *Query {
	q := Query{
		operation:  "",
		options:    nil,
		conditions: make(M),
		field:      "",
	}
	return &q
}

func NewQuery() *Query {
	q := Query{
		operation: "",
		options: &Options{
			skip:       0,
			projection: make(M),
		},
		conditions: make(M),
		field:      "",
	}
	return &q
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

func (q *Query) GetSelect() M {
	return q.options.projection
}

func mappingStringToFieldSets(value Input) Input {
	obj := make(M)
	switch value.(type) {
	case string:
		strArray := strings.Fields(strings.TrimSpace(value.(string)))
		for _, v := range strArray {
			if v[0] == '-' {
				v = v[1:]
				//TODO: Projection 0
				obj[v] = -1
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
	q.options.sort = chain(mappingStringToFieldSets)(value).(M)
	return q
}

func (q *Query) Select(value interface{}) *Query {
	q.options.projection = chain(mappingStringToFieldSets)(value).(M)
	return q
}

func (q *Query) Skip(value int32) *Query {
	q.options.skip = value
	return q
}
func (q *Query) Where(field string) *Query {
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

/*
	e.g. query.Or([]M{{"name": "weny"}, {"age": "20"}})
*/
func (q *Query) Or(value interface{}) *Query {
	flag := "$or"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

/*
	e.g. query.Nor([]M{{"name": "weny"}, {"age": "20"}})
*/
func (q *Query) Nor(value interface{}) *Query {
	flag := "$nor"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) And(value interface{}) *Query {
	flag := "$and"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) Not(value interface{}) *Query {
	flag := "$not"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) Gt(value interface{}) *Query {
	flag := "$gt"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) Gte(value interface{}) *Query {
	flag := "$gte"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) Lt(value interface{}) *Query {
	flag := "$lt"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) Lte(value interface{}) *Query {
	flag := "$lte"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) Ne(value interface{}) *Query {
	flag := "$ne"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) In(value interface{}) *Query {
	flag := "$in"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) Nin(value interface{}) *Query {
	flag := "$nin"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) All(value interface{}) *Query {
	flag := "$all"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) Regex(value interface{}) *Query {
	flag := "$regex"
	if q.hasField() {
		q.bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		)
	} else {
		q.setField(flag).bindCondition(value)
	}
	q.resetField()
	return q
}

func (q *Query) Size(value interface{}) *Query {
	flag := "$size"
	q.ensureField(flag).
		bindCondition(
			chain(
				inputWrapper(flag),
			)(value),
		)
	return q
}

func (q *Query) MaxDistance(value interface{}) *Query {
	flag := "$maxDistance"
	q.ensureField(flag).
		bindCondition(
			chain(
				inputWrapper(flag),
			)(value),
		)
	return q
}

func (q *Query) MinDistance(value interface{}) *Query {
	flag := "$minDistance"
	q.
		ensureField(flag).
		bindCondition(
			chain(
				inputWrapper(flag),
			)(value),
		)
	return q
}

func (q *Query) Mod(value interface{}) *Query {
	flag := "$mod"
	q.
		ensureField(flag).
		bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		).
		resetField()
	return q
}

//TODO: geometry

func (q *Query) Exists(value bool) *Query {
	flag := "$exists"
	q.ensureField(flag).
		bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder)(value),
		).
		resetField()
	return q
}

func (q *Query) ElemMatch(value interface{}) *Query {
	flag := "$elemMatch"
	q.
		ensureField(flag).
		bindCondition(
			chain(
				inputWrapper(flag),
				inputBuilder,
			)(value),
		).
		resetField()
	return q
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
