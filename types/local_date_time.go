package types

import (
	"fmt"
	"strings"
	"time"
)

const StandardDateTimeFormat = "2006-01-02T15:04:05Z"

type LocalDateTime struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *LocalDateTime) UnmarshalJSON(data []byte) error {
	timeStr := string(data)
	timeStr = strings.ReplaceAll(timeStr, `"`, "")

	t, err := time.Parse(StandardDateTimeFormat, timeStr)

	if err != nil {
		return err
	}

	d.Time = time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		time.Local,
	)

	return nil
}

// UnmarshalJSON implements json.Marshaler
func (d LocalDateTime) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%s"`, d.Format(StandardDateTimeFormat)), nil
}

// UnmarshalJSON implements encoding.TextMarshaler
func (d LocalDateTime) MarshalText() (text []byte, err error) {
	return fmt.Appendf(nil, "%s", d.Format(StandardDateTimeFormat)), nil
}

func (d *LocalDateTime) UnmarshalText(text []byte) error {
	timeStr := string(text)

	t, err := time.Parse(StandardDateTimeFormat, timeStr)

	if err != nil {
		return err
	}

	d.Time = time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		time.Local,
	)

	return nil
}

func MustParseLocalDateTime(text string) LocalDateTime {
	t, err := time.Parse(StandardDateTimeFormat, text)

	if err != nil {
		panic(err)
	}

	dateTime := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		time.Local,
	)

	return LocalDateTime{Time: dateTime}
}
