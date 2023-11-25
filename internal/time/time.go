package time

import "time"

func MilliToTime(milli int64) string {
	return time.UnixMilli(milli).UTC().Format("2006-01-02T15:04:05Z")
}
