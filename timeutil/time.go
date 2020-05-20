package timeutil

import (
	"database/sql"
	"time"
)

func FormatCGTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func MaybeFormatCGTime(t sql.NullTime) string {
	if t.Valid {
		return FormatCGTime(t.Time)
	}
	return ""

}

func ParseCGTime(ts string) (time.Time, error) {
	return time.Parse( "2006-01-02 15:04:05", ts)
}

func FormatTimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func Now() string {
	return FormatCGTime(time.Now())
}