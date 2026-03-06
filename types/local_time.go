package types

import (
	"fmt"
	"log/slog"
	"reflect"
	"strings"
	"time"
)

const LocalTimeFormatWithoutZ = "2006-01-02T15:04:05"
const LocalTimeFormat = "2006-01-02T15:04:05Z"

type LocalTime time.Time

func (lt *LocalTime) String() string {
	return time.Time(*lt).Local().Format(LocalTimeFormat)
}

func (lt *LocalTime) Time() time.Time {
	return time.Time(*lt)
}

func (lt LocalTime) MarshalText() (text []byte, err error) {
	return []byte(lt.String()), nil
}

func (lt LocalTime) MarshalJSON() (text []byte, err error) {
	return fmt.Appendf(nil, `"%s"`, lt.String()), nil
}

func (lt *LocalTime) UnmarshalJSON(data []byte) error {
	timeStr := strings.ReplaceAll(string(data), "\"", "")

	slog.Error("FOO", slog.Any("B", timeStr))
	timeFormat := LocalTimeFormatWithoutZ

	if strings.HasSuffix(timeStr, "Z") {
		timeFormat = LocalTimeFormat
	}

	v, err := time.ParseInLocation(timeFormat, timeStr, time.Local)

	if err != nil {
		slog.Error("FOO Err", slog.Any("C", err))
		return err
	}

	*lt = LocalTime(v)

	return nil
}

var LocalTimeConverter = func(value string) (res reflect.Value) {
	defer func() {
		if e := recover(); e != nil {
			res = reflect.Value{}
		}
	}()

	timeFormat := LocalTimeFormatWithoutZ

	if strings.HasSuffix(value, "Z") {
		timeFormat = LocalTimeFormat
	}

	if v, err := time.ParseInLocation(timeFormat, value, time.Local); err == nil {
		return reflect.ValueOf(LocalTime(v))
	}

	return reflect.Value{}
}
