package time

import "time"

func FormatCGTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}