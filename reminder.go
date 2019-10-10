package calendar

import (
	"time"
)

func init() {
	if time.Now().After(time.Date(Version, time.October, 1, 0,0,0,0,time.UTC)) {
		println("WARNING: you should update https://github.com/serjvanilla/calendar for next year")
	}
}
