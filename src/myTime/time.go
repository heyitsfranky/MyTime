package myTime

import (
	"time"
)

func IsDateTimeFormatCorrect(input string) bool {
	layout := "2006-01-02T15:04:05.999Z"
	_, err := time.Parse(layout, input)
	if err != nil {
		return false
	}
	return err == nil
}

func GetCurrentDateTimeWithMilliseconds() string {
	now := time.Now().UTC()
	return now.Format("2006-01-02T15:04:05.999Z")
}

func GetCurrentDateTime() string {
	now := time.Now().UTC()
	return now.Format("2006-01-02T15:04:05Z")
}

func GetCurrentDate() string {
	now := time.Now().UTC()
	return now.Format("2006-01-02T00:00:00Z")
}

func GetYesterdaysDateTime() string {
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	return yesterday.Format("2006-01-02T15:04:05Z")
}

func GetYesterdaysDate() string {
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	return yesterday.Format("2006-01-02T00:00:00Z")
}

func IsBefore(timeStr1, timeStr2 string) (bool, error) {
	t1, err := time.Parse("2006-01-02T15:04:05.999Z", timeStr1)
	if err != nil {
		return false, err
	}
	t2, err := time.Parse("2006-01-02T15:04:05.999Z", timeStr2)
	if err != nil {
		return false, err
	}
	return t1.Before(t2), nil
}

// per ISO8601, a week starts with "Monday" thus Monday is "0"
func GetDay(dateStr string) int {
	layout := "2006-01-02T15:04:05.999Z"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return -1
	}
	dayOfWeek := int(date.Weekday())
	return (dayOfWeek + 6) % 7
}

func GetCalendarWeekAndYear(date string) (int, int, error) {
	layout := "2006-01-02T15:04:05.999Z"
	tn, err := time.Parse(layout, date)
	if err != nil {
		return 0, 0, err
	}
	year, calendarWeek := tn.ISOWeek()
	return year, calendarWeek, nil
}
