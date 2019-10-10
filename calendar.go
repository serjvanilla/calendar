package calendar

import "time"

// IsWorkDay determines a date is a working day or not
func IsWorkDay(date time.Time) bool {
	y, m, d := date.Date()

	if is, ok := reschedule[y][m][d]; ok {
		return is
	} else if _, ok := defaultHolidays[m][d]; ok {
		return false
	} else if wd := date.Weekday(); wd == time.Saturday || wd == time.Sunday {
		return false
	}

	return true
}

// Add return the date corresponding to adding the given number of working days to date
func Add(date time.Time, days int) time.Time {
	var inc int
	if days > 0 {
		inc = 1
	} else {
		inc = -1
	}

	for {
		if days == 0 {
			break
		}
		date = date.AddDate(0, 0, inc)
		if IsWorkDay(date) {
			days -= inc
		}
	}

	return date
}

// Period calculates count of working day between from and to
func Period(from, to time.Time) int {
	var inc int
	if from.Before(to) {
		inc = 1
	} else {
		inc = -1
	}

	var days int
	for {
		if from.Equal(to) {
			break
		}
		from = from.AddDate(0, 0, inc)
		if IsWorkDay(from) {
			days += inc
		}
	}

	return days
}

var defaultHolidays = map[time.Month]map[int]struct{}{
	time.January:  {1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}},
	time.February: {23: {}},
	time.March:    {8: {}},
	time.May:      {1: {}, 9: {}},
	time.June:     {12: {}},
	time.November: {4: {}},
}
