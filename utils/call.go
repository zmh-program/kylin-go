package utils

import (
	"reflect"
)

func CallFunc(_fn interface{}, _args []interface{}) interface{} {
	fn := reflect.ValueOf(_fn)

	args := make([]reflect.Value, len(_args))
	for i, arg := range _args {
		args[i] = reflect.ValueOf(arg)
	}

	resp := fn.Call(args)
	if len(resp) == 0 {
		return nil
	} else if len(resp) == 1 {
		return resp[0].Interface()
	} else {
		results := make([]interface{}, len(resp))
		for i, r := range resp {
			results[i] = r.Interface()
		}
		return results
	}
}
