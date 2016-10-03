package util

import "time"

func FormatDate(t time.Time) string {
	if t.IsZero() {
		return "N/A"
	}
	return t.Format("2006-01-02, 15:04:05")
}
