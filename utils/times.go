package utils

import (
	"strconv"
	"time"
)

//TimeSince returns string for time duration like "3 minutes ago" ...
func TimeSince(s time.Time) string {

	now := time.Now()

	dur := now.Sub(s)
	if dur.Seconds() <= 0 {
		return "future!"
	}

	years := int64(dur.Hours() / 8760.0) 
	if years > 1 {
		return strconv.FormatInt(years, 10) + " years ago"
	}

	months := int64(dur.Hours() / 720.0)
	if months > 1 {
		return strconv.FormatInt(months, 10) + " months ago"
	}

	days := int64(dur.Hours() / 24.0)
	if days > 1 {
		return strconv.FormatInt(days, 10) + " days ago"
	}

	if dur.Hours() > 1 {
		return strconv.FormatInt(int64(dur.Hours()), 10) + " hours ago"
	}

	if dur.Minutes() > 1 {
		return strconv.FormatInt(int64(dur.Minutes()), 10) + " minutes ago"
	}

	return strconv.FormatInt(int64(dur.Seconds()), 10) + " seconds ago"
}
