package gtime

import (
	"log"
	"testing"
	"time"
)

func TestGetNowWeek(t *testing.T) {
	day := GetNowWeek()
	now := time.Now()
	if day != now.Weekday().String() {
		t.Error("GetNowWeek error")
	}
	log.Println(day)
	t.Log(now)
}
