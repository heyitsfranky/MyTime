package MyTime

import (
	"testing"
	"time"
)

func Test_GetCurrentDateTimeWithMilliseconds(t *testing.T) {
	t.Run("ValidFormat", func(t *testing.T) {
		expected := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
		result := GetCurrentDateTimeWithMilliseconds()
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})
	t.Run("NonEmptyResult", func(t *testing.T) {
		result := GetCurrentDateTimeWithMilliseconds()
		if result == "" {
			t.Error("Expected a non-empty result")
		}
	})
}

func Test_GetCurrentDateTime(t *testing.T) {
	t.Run("ValidFormat", func(t *testing.T) {
		expected := time.Now().UTC().Format("2006-01-02T15:04:05Z")
		result := GetCurrentDateTime()
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})
	t.Run("NonEmptyResult", func(t *testing.T) {
		result := GetCurrentDateTime()
		if result == "" {
			t.Error("Expected a non-empty result")
		}
	})
}

func Test_GetCurrentDate(t *testing.T) {
	t.Run("ValidFormat", func(t *testing.T) {
		expected := time.Now().UTC().Format("2006-01-02T00:00:00Z")
		result := GetCurrentDate()
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})
	t.Run("NonEmptyResult", func(t *testing.T) {
		result := GetCurrentDate()
		if result == "" {
			t.Error("Expected a non-empty result")
		}
	})
}

func Test_GetYesterdayTime(t *testing.T) {
	t.Run("ValidFormat", func(t *testing.T) {
		expected := time.Now().UTC().Add(-24 * time.Hour).Format("2006-01-02T15:04:05Z")
		result := GetYesterdaysDateTime()
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})
	t.Run("NonEmptyResult", func(t *testing.T) {
		result := GetYesterdaysDateTime()
		if result == "" {
			t.Error("Expected a non-empty result")
		}
	})
}

func Test_GetYesterdayDate(t *testing.T) {
	t.Run("ValidFormat", func(t *testing.T) {
		expected := time.Now().UTC().Add(-24 * time.Hour).Format("2006-01-02T00:00:00Z")
		result := GetYesterdaysDate()
		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	})
	t.Run("NonEmptyResult", func(t *testing.T) {
		result := GetYesterdaysDate()
		if result == "" {
			t.Error("Expected a non-empty result")
		}
	})
}

func Test_IsBefore(t *testing.T) {
	t.Run("BeforeTimestamp", func(t *testing.T) {
		t1 := "2006-01-02T15:04:05.999Z"
		t2 := "2023-09-16T12:00:00.000Z"
		expected := true
		result, err := IsBefore(t1, t2)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("Expected: %v, Got: %v", expected, result)
		}
	})
	t.Run("AfterTimestamp", func(t *testing.T) {
		t1 := "2023-09-16T12:00:00.000Z"
		t2 := "2006-01-02T15:04:05.999Z"
		expected := false
		result, err := IsBefore(t1, t2)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("Expected: %v, Got: %v", expected, result)
		}
	})
}

func Test_IsDateTimeFormatCorrect(t *testing.T) {
	t.Run("ValidFormat", func(t *testing.T) {
		input := "2023-09-16T12:00:00.000Z"
		expected := true
		result := IsDateTimeFormatCorrect(input)
		if result != expected {
			t.Errorf("Expected: %v, Got: %v", expected, result)
		}
	})
	t.Run("InvalidFormat", func(t *testing.T) {
		input := "2023-09-16 12:00:00"
		expected := false
		result := IsDateTimeFormatCorrect(input)
		if result != expected {
			t.Errorf("Expected: %v, Got: %v", expected, result)
		}
	})
}

func Test_GetCalendarWeekAndYear(t *testing.T) {
	t.Run("ValidDate", func(t *testing.T) {
		date := "2023-09-16T12:00:00.000Z"
		expectedYear := 2023
		expectedWeek := 37
		year, week, err := GetCalendarWeekAndYear(date)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if year != expectedYear || week != expectedWeek {
			t.Errorf("Expected Year: %d, Got: %d; Expected Week: %d, Got: %d", expectedYear, year, expectedWeek, week)
		}
	})
	t.Run("InvalidDate", func(t *testing.T) {
		date := "invalid_date"
		_, _, err := GetCalendarWeekAndYear(date)
		if err == nil {
			t.Error("Expected an error for an invalid date format")
		}
	})
}

func Test_GetDay(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"2023-09-18T15:04:05.999Z", 0}, // Monday
		{"2023-09-19T15:04:05.999Z", 1}, // Tuesday
		{"2023-09-20T15:04:05.999Z", 2}, // Wednesday
		{"2023-09-21T15:04:05.999Z", 3}, // Thursday
		{"2023-09-22T15:04:05.999Z", 4}, // Friday
		{"2023-09-23T15:04:05.999Z", 5}, // Saturday
		{"2023-09-24T15:04:05.999Z", 6}, // Sunday
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := GetDay(test.input)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}

func Test_GetGivenDateTimeInCorrectFormat(t *testing.T) {
	testCases := []struct {
		input    time.Time
		expected string
	}{
		{
			input:    time.Date(2023, time.September, 22, 12, 30, 45, 999000000, time.UTC),
			expected: "2023-09-22T12:30:45.999Z",
		},
		{
			input:    time.Date(1990, time.May, 5, 10, 15, 30, 500000000, time.UTC),
			expected: "1990-05-05T10:15:30.5Z",
		},
		{
			input:    time.Date(2021, time.December, 31, 23, 59, 59, 999999999, time.UTC),
			expected: "2021-12-31T23:59:59.999Z",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.expected, func(t *testing.T) {
			result := GetGivenDateTimeInCorrectFormat(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected %s, but got %s", testCase.expected, result)
			}
		})
	}
}
