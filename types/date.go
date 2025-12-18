package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

const StandardDateFormat = "2006-01-02"

func (a Date) Value() (driver.Value, error) {
	return a.Time.Format(StandardDateFormat), nil
}

func (a *Date) Scan(value interface{}) error {
	if value == nil {

		return nil
	}

	if dateTime, ok := value.(time.Time); ok {
		a.Time = time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, time.UTC)
		return nil
	}

	return nil
}

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

func (d Date) Between(d1 Date, d2 Date) bool {
	return d.Equal(d1.Time) || d.Equal(d2.Time) || (d.After(d1.Time) && d.Before(d2.Time))
}

func MustParseDateLocal(data string) Date {
	res, err := time.Parse(StandardDateFormat, data)

	if err != nil {
		panic(err.Error())
	}

	return Date{Time: time.Date(res.Year(), res.Month(), res.Day(), 0, 0, 0, 0, time.Local)}
}

func DatetimeToDate(data time.Time) Date {
	return Date{
		Time: time.Date(data.Year(), data.Month(), data.Day(), 0, 0, 0, 0, time.UTC),
	}
}
