package types

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

const StandardDateFormat = "2006-01-02"

// UnmarshalJSON implements json.Unmarshaler.
func (d *Date) UnmarshalJSON(data []byte) error {
	timeStr := string(data)
	timeStr = strings.ReplaceAll(timeStr, `"`, "")

	t, err := time.Parse(StandardDateFormat, timeStr)

	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

// UnmarshalJSON implements json.Marshaler
func (d Date) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%s"`, d.Format(StandardDateFormat)), nil
}

// UnmarshalJSON implements encoding.TextMarshaler
func (d Date) MarshalText() (text []byte, err error) {
	return fmt.Appendf(nil, "%s", d.Format(StandardDateFormat)), nil
}

func (d *Date) UnmarshalText(text []byte) error {
	timeStr := string(text)

	t, err := time.Parse(StandardDateFormat, timeStr)

	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

func MustParseDate(text string) Date {
	t, err := time.Parse(StandardDateFormat, text)

	if err != nil {
		panic(err)
	}

	return Date{Time: t}
}
