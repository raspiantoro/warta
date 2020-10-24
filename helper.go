package warta

import (
	"errors"
	"reflect"
)

func compareAndAnalyze(listener interface{}, args []interface{}) error {
	x := reflect.TypeOf(listener)
	in := x.NumIn()

	if in != len(args) {
		return errors.New("Arguments length not match")
	}

	for i := 0; i < in; i++ {
		inV := x.In(i)
		argsV := args[i]

		if inV.Kind() != reflect.ValueOf(argsV).Kind() {
			return errors.New("Call using different kind of arguments")
		}
	}

	return nil
}
