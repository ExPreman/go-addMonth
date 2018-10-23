package main

import (
	"fmt"
	"time"
)

func AddMonth(input time.Time, month int) time.Time {
	sumOfDays := 0
	for i := month; i > 0; i-- {
		thisMonth := time.Date(input.Year(), time.Month(int(input.Month()) + i + 1), 0, 0, 0, 0, 0, time.UTC)
		sumOfDays += thisMonth.Day()
	}

	return input.AddDate(0, 0, sumOfDays)
}

func main() {
	date := time.Date(2018, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(AddMonth(date, 3))
}