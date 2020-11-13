/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 9/3/20 11:55 PM
*/

package godm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestQuery_WithoutWhere(t *testing.T) {
//	//assert.Panics(t, func() {
//	//	NewQuery().Gt(10)
//	//})
//}
func TestQuery_Or(t *testing.T) {
	var query = NewQuery()
	query.Or([]M{{"name": "weny"}, {"age": "20"}})
	assert.Equal(t, []M{{"name": "weny"}, {"age": "20"}}, query.conditions["$or"])
	query.Or([]M{{"age": "22"}})
	assert.Equal(t, []M{{"age": "22"}}, query.conditions["$or"])
}
func TestQuery_And(t *testing.T) {
	var query = NewQuery()
	query.And([]M{{"name": "weny"}, {"age": "20"}})
	assert.Equal(t, []M{{"name": "weny"}, {"age": "20"}}, query.conditions["$and"])
	query.And([]M{{"age": "22"}})
	assert.Equal(t, []M{{"age": "22"}}, query.conditions["$and"])
}
func TestQuery_Nor(t *testing.T) {
	var query = NewQuery()
	query.Nor([]M{{"name": "weny"}, {"age": "20"}})
	assert.Equal(t, []M{{"name": "weny"}, {"age": "20"}}, query.conditions["$nor"])
	query.Nor([]M{{"age": "22"}})
	assert.Equal(t, []M{{"age": "22"}}, query.conditions["$nor"])
}

func TestQuery_Where(t *testing.T) {
	var query = NewQuery()
	query.Where("name").Eq("weny")
	assert.Equal(t, M{"name": "weny"}, query.conditions)
	query2 := NewQuery().
		Where("user").
		Eq(func(query *Query) {
			query.
				Where("name").
				Eq("weny").
				Where("gender").
				Eq("male")
		}).
		Where("data").
		Eq(func(query *Query) {
			query.
				Where("version").
				Gt(10).
				Where("tags").
				ElemMatch(func(q *Query) {
					q.In(A{1, 2, 3, 4})
				})
		})
	assert.Equal(t,
		M{
			"user": M{"name": "weny", "gender": "male"},
			"data": M{
				"version": M{
					"$gt": 10,
				},
				"tags": M{
					"$elemMatch": M{
						"$in": A{1, 2, 3, 4},
					},
				},
			},
		}, query2.conditions)
}

func TestQuery_Select(t *testing.T) {
	query := NewQuery()
	query.Select("name -_id")
	assert.Equal(t, M{"name": 1, "_id": -1}, query.options.projection)

	query2 := NewQuery()
	query2.Select(M{"name": 1, "_id": -1})
	assert.Equal(t, M{"name": 1, "_id": -1}, query2.options.projection)
}

func TestQuery_Sort(t *testing.T) {
	query := NewQuery()
	query.Sort("name -_id")
	assert.Equal(t, M{"name": 1, "_id": -1}, query.options.sort)

	query2 := NewQuery()
	query2.Sort(M{"name": 1, "_id": -1})
	assert.Equal(t, M{"name": 1, "_id": -1}, query2.options.sort)
}

func TestQuery_ElemMatch(t *testing.T) {
	query := NewQuery()
	query.
		Where("tags").
		ElemMatch(func(q *Query) {
			q.
				Where("version").
				Eq(1).
				Where("env").
				Eq("dev")
		})

	assert.Equal(t, M{
		"tags": M{
			"$elemMatch": M{
				"version": 1,
				"env":     "dev",
			},
		},
	}, query.conditions)

	query3 := NewQuery()
	query3.
		Where("tags").
		ElemMatch(func(q *Query) *Query {
			return q.
				Where("version").
				Eq(1).
				Where("env").
				Eq("dev")

		})

	assert.Equal(t, M{
		"tags": M{
			"$elemMatch": M{
				"version": 1,
				"env":     "dev",
			},
		},
	}, query3.conditions)

	query2 := NewQuery()
	query2.
		Where("versions").
		ElemMatch([]uint{1, 2, 3})
	assert.Equal(t, M{
		"versions": M{
			"$elemMatch": []uint{1, 2, 3},
		},
	}, query2.conditions)
}

type ConditionBuilderTest struct {
	fn                    interface{}
	flag                  string
	field                 string
	input                 interface{}
	expectedResultBuilder func(field string, flag string, input interface{}) (res M)
}

func TestQuery_TestConditionBuilder(t *testing.T) {
	expectedResultBuilder := func(field string, flag string, input interface{}) (res M) {
		res = make(M)
		res[field] = M{flag: input}
		return
	}
	testSets := []ConditionBuilderTest{
		{
			fn:                    NewQuery().Where("test").Gt,
			input:                 10,
			field:                 "test",
			flag:                  "$gt",
			expectedResultBuilder: expectedResultBuilder,
		},
		{
			fn:                    NewQuery().Where("test").Gte,
			input:                 10,
			field:                 "test",
			flag:                  "$gte",
			expectedResultBuilder: expectedResultBuilder,
		},
		{
			fn:                    NewQuery().Where("test").Lt,
			input:                 10,
			field:                 "test",
			flag:                  "$lt",
			expectedResultBuilder: expectedResultBuilder,
		},
		{
			fn:                    NewQuery().Where("test").Lte,
			input:                 10,
			field:                 "test",
			flag:                  "$lte",
			expectedResultBuilder: expectedResultBuilder,
		},
		{
			fn:                    NewQuery().Where("test").Ne,
			input:                 10,
			field:                 "test",
			flag:                  "$ne",
			expectedResultBuilder: expectedResultBuilder,
		},
		{
			fn:                    NewQuery().Where("test").In,
			input:                 A{1, 2, 3, 4},
			field:                 "test",
			flag:                  "$in",
			expectedResultBuilder: expectedResultBuilder,
		},
		{
			fn:                    NewQuery().Where("test").Nin,
			input:                 A{1, 2, 3, 4},
			field:                 "test",
			flag:                  "$nin",
			expectedResultBuilder: expectedResultBuilder,
		},
		{
			fn:                    NewQuery().Where("test").All,
			input:                 A{1, 2, 3, 4},
			field:                 "test",
			flag:                  "$all",
			expectedResultBuilder: expectedResultBuilder,
		},
		{
			fn:                    NewQuery().Where("test").Mod,
			input:                 A{1, 2, 3, 4},
			field:                 "test",
			flag:                  "$mod",
			expectedResultBuilder: expectedResultBuilder,
		},
		{
			fn:                    NewQuery().Where("test").Exists,
			input:                 true,
			field:                 "test",
			flag:                  "$exists",
			expectedResultBuilder: expectedResultBuilder,
		},
	}
	for _, v := range testSets {
		var res interface{}
		switch v.fn.(type) {
		case func(interface{}) *Query:
			res = v.fn.(func(interface{}) *Query)(v.input).conditions
			break
		case func(A) *Query:
			res = v.fn.(func(A) *Query)(v.input.(A)).conditions
			break
		case func(bool) *Query:
			res = v.fn.(func(bool) *Query)(v.input.(bool)).conditions
		}
		exp := v.expectedResultBuilder(v.field, v.flag, v.input)
		assert.Equal(t, exp, res)
	}
}
