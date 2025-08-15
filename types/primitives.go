package types

import "reflect"

func IsPrimitive(data any) bool {
	kind := reflect.TypeOf(data).Kind()

	return kind == reflect.Float64 ||
		kind == reflect.Float32 ||
		kind == reflect.String ||
		kind == reflect.Int ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Bool
}
