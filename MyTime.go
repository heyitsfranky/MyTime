package MyTime

import (
	"time"
)

func GetCurrentDateTime() int64 {
	return time.Now().UTC().UnixMilli()
}

func GetGivenDateTime(givenDateTime time.Time) int64 {
	return givenDateTime.UnixMilli()
}

func GetCurrentDate() int64 {
	now := time.Now().UTC()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return midnight.UnixMilli()
}

func GetYesterdaysDateTime() int64 {
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	return yesterday.UnixMilli()
}

func GetYesterdaysDate() int64 {
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	midnight := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.UTC)
	return midnight.UnixMilli()
}

func IsBefore(time1, time2 int64) bool {
	t1 := time.UnixMilli(time1)
	t2 := time.UnixMilli(time2)
	return t1.Before(t2)
}

// per ISO8601, a week starts with "Monday" thus Monday is "0"
func GetDay(input int64) int {
	t := time.UnixMilli(input)
	dayOfWeek := int(t.Weekday())
	return (dayOfWeek + 6) % 7
}

func IsNewWeekTomorrow() bool {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	return tomorrow.Weekday() == time.Monday
}

func GetCalendarWeekAndYear(input int64) (int, int, error) {
	t := time.UnixMilli(input)
	year, calendarWeek := t.ISOWeek()
	return year, calendarWeek, nil
}
