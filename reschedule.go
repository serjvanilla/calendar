package calendar

import "time"

const Version  = 2020

var reschedule = map[int]map[time.Month]map[int]bool{
	2018: {
		time.March: {
			9: false,
		},
		time.April: {
			28: true,
			30: false,
		},
		time.May: {
			2: false,
		},
		time.June: {
			9:  true,
			11: false,
		},
		time.December: {
			29: true,
			31: false,
		},
	},
	2019: {
		time.May: {
			2:  false,
			3:  false,
			10: false,
		},
	},
	2020: {
		time.May: {
			4: false,
			5: false,
		},
	},
}
