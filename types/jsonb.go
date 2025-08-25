package types

import (
	"database/sql/driver"
	"encoding/json"
)

type JSONB map[string]any

// Scan implements sql.Scanner.
func (j *JSONB) Scan(src any) error {
	var result JSONB

	if data, ok := src.([]byte); ok {
		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}
	}

	if data, ok := src.(string); ok {
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			return err
		}
	}

	*j = result

	return nil
}

// Value implements driver.Valuer.
func (j JSONB) Value() (driver.Value, error) {
	if len(j) == 0 {
		return []byte(nil), nil
	}

	res, err := json.Marshal(j)

	return res, err
}

func ToJSONB(data any) (JSONB, error) {
	var result JSONB
	jsonByte, err := json.Marshal(data)

	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(jsonByte, &result); err != nil {
		return result, err
	}

	return result, nil
}
