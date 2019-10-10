package calendar

import (
	"reflect"
	"testing"
	"time"
)

func TestIsWorkDay(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "rescheduled_1",
			args: args{
				date: time.Date(2018, 4, 28, 0, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "rescheduled_2",
			args: args{
				date: time.Date(2018, 4, 30, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "default_1",
			args: args{
				date: time.Date(2019, 11, 4, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "off_day_1",
			args: args{
				date: time.Date(2019, 10, 12, 0, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "workday",
			args: args{
				date: time.Date(2019, 10, 8, 0, 0, 0, 0, time.UTC),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWorkDay(tt.args.date); got != tt.want {
				t.Errorf("IsWorkDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		date time.Time
		days int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "common",
			args: args{
				date: time.Date(2019, 10, 8, 0, 0, 0, 0, time.UTC),
				days: 15,
			},
			want: time.Date(2019, 10, 29, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "negative",
			args: args{
				date: time.Date(2019, 10, 29, 0, 0, 0, 0, time.UTC),
				days: -15,
			},
			want: time.Date(2019, 10, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "new year",
			args: args{
				date: time.Date(2019, 12, 30, 0, 0, 0, 0, time.UTC),
				days: 15,
			},
			want: time.Date(2020, 1, 28, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.date, tt.args.days); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeriod(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive",
			args: args{
				from: time.Date(2019, 10, 8, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 10, 29, 0, 0, 0, 0, time.UTC),
			},
			want: 15,
		},
		{
			name: "negative",
			args: args{
				from: time.Date(2019, 10, 29, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 10, 8, 0, 0, 0, 0, time.UTC),
			},
			want: -15,
		},
		{
			name: "whole year",
			args: args{
				from: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
				to:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: 247,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Period(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("Period() = %v, want %v", got, tt.want)
			}
		})
	}
}
