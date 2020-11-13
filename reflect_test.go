/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 9/5/20 1:03 PM
*/

package godm

import (
	"fmt"
	"reflect"
	"testing"
)

type TestModel struct {
	id   string
	name string
}

func TestReflect(t *testing.T) {
	m := TestModel{
		id:   "test",
		name: "name",
	}
	fmt.Println("TypeOf :", reflect.TypeOf(m))

}
