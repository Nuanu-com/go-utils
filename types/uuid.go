package types

import (
	"reflect"

	"github.com/google/uuid"
)

var UUIDConverter = func(value string) reflect.Value {
	if res, err := uuid.Parse(value); err == nil {
		return reflect.ValueOf(res)
	}
	return reflect.Value{}
}
