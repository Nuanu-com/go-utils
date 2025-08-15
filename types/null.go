package types

import (
	"bytes"
	"database/sql"
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/Nuanu-com/go-utils/slice_utils"

	"github.com/google/uuid"
)

type Null[T any] sql.Null[T]

// UnmarshalJSON implements json.Unmarshaler.
func (n *Null[T]) UnmarshalJSON(data []byte) error {
	var result T

	if data == nil {
		return nil
	}

	if bytes.Equal(data, []byte(`null`)) {
		return nil
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	n.V = result
	n.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler.
func (n Null[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}

	data, err := json.Marshal(n.V)

	return data, err
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *Null[T]) UnmarshalText(data []byte) error {
	if data == nil {
		return nil
	}

	if bytes.Equal(data, []byte(`null`)) {
		return nil
	}

	dataT := reflect.ValueOf(&n.V).Elem()

	// TODO: please improve slice handling
	if dataT.Type().Kind() == reflect.Slice {
		itemType := reflect.TypeOf(n.V).Elem().Kind()
		dataStr := string(data)

		if itemType == reflect.String ||
			itemType == reflect.TypeFor[uuid.UUID]().Kind() ||
			itemType == reflect.TypeFor[Date]().Kind() {
			parts := strings.Split(dataStr, ",")
			parts = slice_utils.Map(parts, func(part string) string { return fmt.Sprintf(`"%s"`, part) })
			dataStr = strings.Join(parts, ",")
		}

		dataStr = fmt.Sprintf(`[%s]`, dataStr)

		var result T

		if err := json.Unmarshal([]byte(dataStr), &result); err != nil {
			return err
		}

		dataT.Set(reflect.ValueOf(result))
		return nil
	}

	v, err := unmarshalSingleValue[T](data)

	if err != nil {
		return err
	}

	n.V = v
	n.Valid = true
	return nil
}

func unmarshalSingleValue[T any](data []byte) (T, error) {
	var result T
	if reflect.TypeFor[T]().Kind() == reflect.String {
		v := strings.TrimSpace(string(data))
		v = fmt.Sprintf(`"%s"`, v)

		if err := json.Unmarshal([]byte(v), &result); err == nil {
			return result, nil
		}
	}

	if IsPrimitive(result) {
		if err := json.Unmarshal(data, &result); err == nil {
			return result, nil
		}
	}

	if un, ok := any(&result).(encoding.TextUnmarshaler); ok {
		if err := un.UnmarshalText(data); err == nil {
			return result, nil
		}
	}

	return result, fmt.Errorf("Please implement encoding.TextUnmarshaler for type %v", reflect.TypeFor[T]().String())
}

// MarshalText implements encoding.TextMarshaler.
func (n Null[T]) MarshalText() (text []byte, err error) {
	if !n.Valid {
		return nil, nil
	}

	if m, ok := any(n.V).(encoding.TextMarshaler); ok {
		return m.MarshalText()
	}

	if s, ok := any(n.V).(string); ok {
		return []byte(s), nil
	}

	dataKind := reflect.TypeOf(n.V).Kind()

	if dataKind == reflect.Int ||
		dataKind == reflect.Float32 ||
		dataKind == reflect.Float64 {
		return fmt.Appendf(nil, "%d", any(n.V)), nil
	}

	if IsPrimitive(any(n.V)) {
		return fmt.Appendf(nil, "%v", any(n.V)), nil
	}

	dataT := reflect.ValueOf(&n.V).Elem()

	if dataT.Type().Kind() == reflect.Slice {
		dataByte, err := json.Marshal(n.V)

		if err != nil {
			return nil, err
		}

		if dataByte == nil || bytes.Equal(dataByte, []byte(`null`)) {
			return nil, nil
		}

		dataStr := string(dataByte)

		dataStr = strings.ReplaceAll(dataStr, "[", "")
		dataStr = strings.ReplaceAll(dataStr, "]", "")
		dataStr = strings.ReplaceAll(dataStr, `"`, "")

		return []byte(dataStr), nil
	}

	return nil, fmt.Errorf("Please implement encoding.TextMarshaler for type %v", reflect.TypeFor[T]().String())
}
