/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 9/4/20 1:17 PM
*/

package godm

type Input interface{}

type InputEndpoint func(Input) Input

func chain(outer InputEndpoint, others ...InputEndpoint) InputEndpoint {
	return func(next Input) Input {
		for i := len(others) - 1; i >= 0; i-- { // reverse
			next = others[i](next)
		}
		return outer(next)
	}
}
