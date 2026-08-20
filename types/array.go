package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONBArray[T any] []T

// Value Marshal
func (a JSONBArray[T]) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONBArray[T]) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func NewJSONBArray[T any](data []T) JSONBArray[T] {
	return data
}
