package date_utils

import "time"

const (
	stringDateLayout   = "2006-01-02T15:04:05Z"
	stringDBDateLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(stringDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(stringDBDateLayout)
}
