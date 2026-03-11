package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const HourMinuteSecondLayout = "15:04:05"
const HourMinuteLayout = "15:04"

type TimeOnly struct {
	time.Time
}

// Value Marshal
func (a TimeOnly) Value() (driver.Value, error) {
	return a.Time.Format(HourMinuteSecondLayout), nil
}

func (a *TimeOnly) Scan(value any) error {
	if value == nil {

		return nil
	}

	if str, ok := value.(string); ok {

		result, err := time.Parse(HourMinuteSecondLayout, str)

		if err != nil {
			return err
		}

		a.Time = result

		a.Time = time.Date(1, 1, 1, result.Hour(), result.Minute(), result.Second(), 0, time.UTC)

		return nil
	}

	return nil
}

func NewTimeOnly(hour int, minute int, second int) TimeOnly {
	return TimeOnly{time.Date(1, 1, 1, hour, minute, second, 0, time.UTC)}
}

func (t *TimeOnly) UnmarshalJSON(data []byte) error {
	timeStr := strings.ReplaceAll(string(data), "\"", "")
	result, err := time.Parse(HourMinuteSecondLayout, timeStr)

	if err != nil {
		return err
	}

	t.Time = time.Date(1, 1, 1, result.Hour(), result.Minute(), result.Second(), 0, time.UTC)
	return nil
}

func (t *TimeOnly) UnmarshalText(data []byte) error {
	timeStr := strings.ReplaceAll(string(data), "\"", "")
	result, err := time.Parse(HourMinuteSecondLayout, timeStr)

	if err != nil {
		return err
	}

	t.Time = time.Date(1, 1, 1, result.Hour(), result.Minute(), result.Second(), 0, time.UTC)
	return nil
}

func (t TimeOnly) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "\"%s\"", t.Format(HourMinuteSecondLayout)), nil
}

func MustParseTimeOnly(data string) TimeOnly {
	res, err := time.Parse(HourMinuteSecondLayout, data)

	if err != nil {
		panic(err.Error())
	}

	return TimeOnly{Time: time.Date(1, 1, 1, res.Hour(), res.Minute(), res.Second(), 0, time.UTC)}
}

func DatetimeToTimeOnly(data time.Time) TimeOnly {
	return TimeOnly{
		Time: time.Date(1, 1, 1, data.Hour(), data.Minute(), data.Second(), 0, time.UTC),
	}
}
