package calendar

import (
	"testing"
)

// Test data structure for isLeapYear
type leapYearTestData struct {
	year     int
	expected bool
}

// Test isLeapYear function
func TestIsLeapYear(t *testing.T) {
	testData := []leapYearTestData{
		{year: 1900, expected: false},
		{year: 4, expected: true},
		{year: 1904, expected: true},
		{year: 2000, expected: true},
	}

	for _, data := range testData {
		if out := isLeapYear(data.year); data.expected != out {
			t.Errorf("Expected isLeapYear(%#v) = %v but got isLeapYear(%#v) = %v instead\n", data.year, data.expected, data.year, out)
		}
	}
}

// Test data structure for getNoOfDays
type noOfDaysTestData struct {
	year  int
	month int
	days  int
}

// Test getNoOfDays function
func TestGetDaysInMonth(t *testing.T) {
	testData := []noOfDaysTestData{
		{year: 1900, month: 1, days: 31},
		{year: 1900, month: 2, days: 28},
		{year: 1900, month: 3, days: 31},
		{year: 1900, month: 4, days: 30},
		{year: 1900, month: 5, days: 31},
		{year: 1900, month: 6, days: 30},
		{year: 1900, month: 7, days: 31},
		{year: 1900, month: 8, days: 31},
		{year: 1900, month: 9, days: 30},
		{year: 1900, month: 10, days: 31},
		{year: 1900, month: 11, days: 30},
		{year: 1900, month: 12, days: 31},
		{year: 1904, month: 1, days: 31},
		{year: 1904, month: 2, days: 29},
		{year: 1904, month: 3, days: 31},
		{year: 1904, month: 4, days: 30},
		{year: 1904, month: 5, days: 31},
		{year: 1904, month: 6, days: 30},
		{year: 1904, month: 7, days: 31},
		{year: 1904, month: 8, days: 31},
		{year: 1904, month: 9, days: 30},
		{year: 1904, month: 10, days: 31},
		{year: 1904, month: 11, days: 30},
		{year: 1904, month: 12, days: 31},
		{year: 2000, month: 1, days: 31},
		{year: 2000, month: 2, days: 29},
		{year: 2000, month: 3, days: 31},
		{year: 2000, month: 4, days: 30},
		{year: 2000, month: 5, days: 31},
		{year: 2000, month: 6, days: 30},
		{year: 2000, month: 7, days: 31},
		{year: 2000, month: 8, days: 31},
		{year: 2000, month: 9, days: 30},
		{year: 2000, month: 10, days: 31},
		{year: 2000, month: 11, days: 30},
		{year: 2000, month: 12, days: 31},
	}

	for _, data := range testData {
		if out := getDaysInMonth(data.year, data.month); data.days != out {
			t.Errorf("Expected getNoOfDays(%v,%v) = %v but got getNoOfDays(%v,%v) = %v instead.\n", data.year, data.month, data.days, data.year, data.month, out)
		}
	}
}

// Test data structure for getMonthStartDayIndex
type monthStartDay struct {
	year     int
	month    int
	dayIndex int
}

func TestGetMonthStartDayIndex(t *testing.T) {
	testData := []monthStartDay{
		{year: 1, month: 1, dayIndex: 1},
		{year: 2, month: 1, dayIndex: 2},
		{year: 5, month: 1, dayIndex: 6},
		{year: 2020, month: 9, dayIndex: 2},
		{year: 2020, month: 6, dayIndex: 1},
		{year: 2022, month: 1, dayIndex: 6},
		{year: 2022, month: 2, dayIndex: 2},
		{year: 2022, month: 3, dayIndex: 2},
		{year: 2022, month: 4, dayIndex: 5},
		{year: 2022, month: 5, dayIndex: 0},
		{year: 2022, month: 6, dayIndex: 3},
		{year: 2022, month: 7, dayIndex: 5},
		{year: 2022, month: 8, dayIndex: 1},
		{year: 2022, month: 9, dayIndex: 4},
		{year: 2022, month: 10, dayIndex: 6},
		{year: 2022, month: 11, dayIndex: 2},
		{year: 2022, month: 12, dayIndex: 4},
	}

	for _, data := range testData {
		if out := getMonthStartDayIndex(data.year, data.month); data.dayIndex != out {
			t.Errorf("Expected getMonthStartDayIndex(%v,%v) = %v got %v instead.\n", data.year, data.month, data.dayIndex, out)
		}
	}

}
