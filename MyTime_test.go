package MyTime

import (
	"fmt"
	"testing"
)

func Test_IsBefore(t *testing.T) {
	time1 := GetCurrentDateTime()
	time2 := GetCurrentDateTime() + 3600000 // Adding an hour to the current time
	isBefore := IsBefore(time1, time2)
	if !isBefore {
		t.Fatalf("Expected %d to be before %d", time1, time2)
	}
}

func TestGetCalendarWeekAndYear(t *testing.T) {
	inputTime := GetCurrentDateTime()
	year, calendarWeek, err := GetCalendarWeekAndYear(inputTime)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Year: %d, Calendar Week: %d\n", year, calendarWeek)
}
