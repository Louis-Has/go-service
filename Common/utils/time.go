package utils

import (
	"time"
)

var timeLocation *time.Location

func init() {
	var err error
	timeLocation, err = time.LoadLocation("Asia/shanghai")
	if err != nil {
		panic(err)
	}
}

func GetCurrentTime() time.Time {
	return time.Now().In(timeLocation)
}

func GetAddTime() time.Time {
	return GetCurrentTime().AddDate(0, 0, 1)
}

func GetFirstDay() time.Time {
	year, month, _ := GetCurrentTime().Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, timeLocation)
}

func GetLastDay() time.Time {
	year, month, _ := GetCurrentTime().Date()
	return time.Date(year, month+1, 1, 0, 0, 0, 0, timeLocation)
}

func CurrentMonthDays() float64 {
	return GetLastDay().Sub(GetFirstDay()).Hours() / 24
}

func GetMonthDays(m int) (int, float64) {
	startTime := GetFirstDay().AddDate(0, m, 0)
	endTime := GetLastDay().AddDate(0, m, 0)
	return int(startTime.Month()), endTime.Sub(startTime).Hours() / 24
}
