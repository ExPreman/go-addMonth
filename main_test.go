package main

import (
	"testing"
	"time"
)

func TestFirstOfDayInMonth(t *testing.T) {
	date := time.Date(2018, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	newDate := AddMonth(date, 3)
	if !newDate.Equal(time.Date(2018,3,31,0,0,0,0, time.UTC)) {
		t.Errorf("date was incorrect")
	}
}

func TestMiddleOfDayInMonth(t *testing.T) {
	date := time.Date(2018, time.Month(1), 15, 0, 0, 0, 0, time.UTC)
	newDate := AddMonth(date, 3)
	if !newDate.Equal(time.Date(2018,4,14,0,0,0,0, time.UTC)) {
		t.Errorf("date was incorrect")
	}
}

func TestEndOfDayInMonth(t *testing.T) {
	date := time.Date(2018, time.Month(1), 31, 0, 0, 0, 0, time.UTC)
	newDate := AddMonth(date, 3)
	if !newDate.Equal(time.Date(2018,4,30,0,0,0,0, time.UTC)) {
		t.Errorf("date was incorrect")
	}
}