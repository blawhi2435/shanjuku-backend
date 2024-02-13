package util

import "time"

const (
	Local = "Local"
)

func ParseDateTime(myTime string) time.Time {
	loc, _ := time.LoadLocation(Local)
	parseTime, _ := time.ParseInLocation("2006/01/02 15:04:05", myTime, loc)
	return parseTime
}