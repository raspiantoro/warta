package warta

import (
	"reflect"
)

func compareAndAnalyze(listener interface{}, args []interface{}) error {
	x := reflect.TypeOf(listener)
	in := x.NumIn()

	if in != len(args) {
		return ErrArgsLenNotMatch
	}

	for i := 0; i < in; i++ {
		inV := x.In(i)
		argsV := args[i]

		if inV.Kind() == reflect.Interface {
			break
		}

		if inV.Kind() != reflect.ValueOf(argsV).Kind() {
			return ErrArgsIsDifferent
		}
	}

	return nil
}
