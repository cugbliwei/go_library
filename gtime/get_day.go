package gtime

import "time"

func GetNowWeek() string {
	now := time.Now()
	return now.Weekday().String()
}

func IsWorkDay() bool {
	now := time.Now()
	if now.Weekday() != time.Saturday && now.Weekday() != time.Sunday {
		return true
	}
	return false
}

func IsWeekDay() bool {
	now := time.Now()
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return true
	}
	return false
}
