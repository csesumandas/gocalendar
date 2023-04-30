// Package calendar has functions and methods to populate a calendar month or year
package calendar

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Year struct {
	year   int        // year value
	months [12]*Month // months inside the year
}

type Month struct {
	year          int
	monthName         // Embedded monthName structure
	noOfDays      int // No of days in the month
	startDayIndex int // Startday index of the month 0 - Sun, 1 - Mon etc
}

type monthName struct {
	mName   string // Month Name like "January", "February" etc
	mSName  string // Month Short Name like "Jan", "Feb" etc
	mNumber int
}

// Create a const array of 12 months and initialize them
var months = [12]monthName{
	{mName: "January", mSName: "Jan", mNumber: 1},
	{mName: "February", mSName: "Feb", mNumber: 2},
	{mName: "March", mSName: "Mar", mNumber: 3},
	{mName: "April", mSName: "Apr", mNumber: 4},
	{mName: "May", mSName: "May", mNumber: 5},
	{mName: "June", mSName: "Jun", mNumber: 6},
	{mName: "July", mSName: "Jul", mNumber: 7},
	{mName: "August", mSName: "Aug", mNumber: 8},
	{mName: "September", mSName: "Sep", mNumber: 9},
	{mName: "October", mSName: "Oct", mNumber: 10},
	{mName: "November", mSName: "Nov", mNumber: 11},
	{mName: "December", mSName: "Dec", mNumber: 12},
}

type weekDayName struct {
	dName  string // Week Day Name like "Sunday", "Monday" etc
	dSName string // Week Day Short Name like "S","M" etc
}

// Create a const array of 7 weekdays and initialize them
var weekDays = [7]weekDayName{
	{dName: "Sunday", dSName: "Su"},
	{dName: "Monday", dSName: "Mo"},
	{dName: "Tuesday", dSName: "Tu"},
	{dName: "Wednesday", dSName: "We"},
	{dName: "Thursday", dSName: "Th"},
	{dName: "Friday", dSName: "Fr"},
	{dName: "Satday", dSName: "Sa"},
}

// WeekDayIndex of 1st Jan of Year 1 in Gregorian Calendar Monday
// According to WeekDays Array the index will be 1
const gregCalFirstDayIndex = 1

// Constant for highlighting current day in the calendar
var invert = "\033[1;30;47m"
var reset = "\033[0m"

// Create a New Year based on the parameters and return the Year
func NewYear(year int) (*Year, error) {
	// Year should be a valid year
	if year < 1 || year > 9999 {
		return nil, fmt.Errorf("Year should be in between 1 and 9999.\n")
	}

	y := Year{}
	y.year = year

	for i := 1; i <= 12; i++ {
		m, err := NewMonth(y.year, i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		y.months[i-1] = m

	}

	return &y, nil
}

// GetMonth gets year and month parameters and returns a Month struct
func NewMonth(year int, month int) (*Month, error) {
	// Year and month should be valid
	if year < 1 || year > 9999 {
		return nil, fmt.Errorf("Year should be in between 1 and 9999.\n")
	} else if month < 1 || month > 12 {
		return nil, fmt.Errorf("Month should be in between 1 and 12.\n")
	}

	m := Month{}
	m.year = year
	m.monthName = months[month-1]
	m.noOfDays = getDaysInMonth(year, month)
	m.startDayIndex = getMonthStartDayIndex(year, month)
	return &m, nil
}

// Checks if the parameter year is leap year or not
func isLeapYear(year int) bool {
	if year%100 != 0 && year%4 == 0 {
		return true
	} else if year%100 == 0 && year%400 == 0 {
		return true
	} else {
		return false
	}
}

// Get No Of Days in the parameter year and month combination
func getDaysInMonth(year int, month int) int {
	noOfDays := 0
	if month <= 7 {
		if month%2 == 0 && month != 2 {
			noOfDays = 30
		} else if month == 2 && isLeapYear(year) {
			noOfDays = 29
		} else if month == 2 && !isLeapYear(year) {
			noOfDays = 28
		} else {
			noOfDays = 31
		}
	} else {
		if month%2 == 0 {
			noOfDays = 31
		} else {
			noOfDays = 30
		}
	}
	return noOfDays
}

// Get Month Start Day Index for the year and month parameter
func getMonthStartDayIndex(year int, month int) int {
	// Calulate no of LeapYears till the prev year
	// For each leap year the days are shifted by 2 due to the fact 366 % 7 = 2
	// For normal years days are shifted by 1 due to the fact 365 % 7 = 1
	prevYear := year - 1
	noOfLeapYears := prevYear/4 - prevYear/100 + prevYear/400

	// Calculate no of days in the year till the start of input month
	// For each pair of months days are 30 + 31 = 61 and this is multiplied by qPrevMonth
	// For each remained month days are added either 30 or 31 based on the month number
	// Based on Leap Year the February will have either 28 or 29 days so either 2 or 1
	// should be deducted to get total no oof days in the year till that month
	prevMonth := month - 1
	qPrevMonth := prevMonth / 2
	rPrevMonth := prevMonth % 2
	noOfDaysTillMonthStart := 0

	if prevMonth <= 7 {
		days := 31
		leapYearDeduct := 0
		if rPrevMonth == 0 {
			days = 30
		}
		if !isLeapYear(year) && qPrevMonth >= 1 {
			leapYearDeduct = 2
		} else if isLeapYear(year) && qPrevMonth >= 1 {
			leapYearDeduct = 1
		}
		noOfDaysTillMonthStart = qPrevMonth*61 + rPrevMonth*days - leapYearDeduct
	} else {
		days := 31
		leapYearDeduct := 0
		if rPrevMonth != 0 {
			days = 30
		}
		if !isLeapYear(year) {
			leapYearDeduct = 2
		} else if isLeapYear(year) {
			leapYearDeduct = 1
		}
		// Added 1 because August has 31 days
		noOfDaysTillMonthStart = qPrevMonth*61 + 1 + rPrevMonth*days - leapYearDeduct
	}

	// Calculate Week Day Index on the start day of input parameter year
	// Formula used is (Days Shift Leap Years
	//					+ Days Shift For All Years
	// 					+ Gregorian Calendar First Day Index
	//					+ No Of Days In The Year Till Month Start
	//					) % 7
	monthStartDayIndex := (noOfLeapYears + (year - 1) + gregCalFirstDayIndex + noOfDaysTillMonthStart) % 7

	return monthStartDayIndex
}

// Print method prints the calendar month in a tabular format similar to linux
// cal command format
func (m *Month) Print(currTime time.Time) {
	// If the system is windows then donot highlight
	if runtime.GOOS == "windows" {
		invert = ""
		reset = ""
	}

	// Get Today from currDate
	currYear := currTime.Year()
	currMonth := int(currTime.Month())
	currDay := currTime.Day()

	// Print Header
	header := fmt.Sprintf("%3s %04d", m.mName, m.year)
	headerPrefixLen := (20 - len(header)) / 2
	format := fmt.Sprintf("%s%ds", "%", headerPrefixLen)
	fmt.Printf(format, " ")
	fmt.Printf("%s\n", header)

	// Print the week days
	for _, weekDay := range weekDays {
		fmt.Printf("%2s ", weekDay.dSName)
	}

	fmt.Println()

	for j := 1; j <= m.startDayIndex; j++ {
		fmt.Printf("   ")
	}

	rowNumber := 1
	// Print the dates
	for i := 1; i <= m.noOfDays; i++ {
		dayPosition := m.startDayIndex + i

		// Highlight the Current Day
		if currYear == m.year && currMonth == m.mNumber && currDay == i {
			fmt.Printf("%s%2d%s ", invert, i, reset)
		} else {
			fmt.Printf("%2d ", i)
		}

		if dayPosition >= 7*rowNumber {
			fmt.Println()
			rowNumber++
		}
	}

	// Send terminal prompt to next line
	fmt.Println()
}

// PrintCalendar function prints the calendar for a whole year in the terminal
// in 4 * 3 grids assuming default terminal width is 80 columns
func (y *Year) Print(currTime time.Time) {
	// If the system is windows then donot highlight
	if runtime.GOOS == "windows" {
		invert = ""
		reset = ""
	}

	// Get Today from currDate
	currYear := currTime.Year()
	currMonth := int(currTime.Month())
	currDay := currTime.Day()

	// Print the Header. PrefixLen is calculated to center the Header
	yearHeader := fmt.Sprintf("%04d", y.year)
	yearHeaderPrefixLen := (66 - 4) / 2
	headerPrefixFormat := fmt.Sprintf("%s%ds", "%", yearHeaderPrefixLen)
	fmt.Printf(headerPrefixFormat, " ")
	fmt.Printf("%s", yearHeader)
	fmt.Println()

	// Get all the months
	yearMonths := y.months

	// Print the Calendar in a 4 * 3 matrix
	for row := 1; row <= 4; row++ {

		// Print the Month Names
		for col := 1; col <= 3; col++ {

			// Calculate month index based on row and col values
			mNo := col + (row-1)*3
			m := yearMonths[mNo-1]

			// Print the Month Name Headers
			monthHeader := fmt.Sprintf("%s", m.mName)
			monthHeaderPrefixLen := (20-len(monthHeader))/2 + (20-len(monthHeader))%2
			monthHeaderPrefixFormat := fmt.Sprintf("%s%ds", "%", monthHeaderPrefixLen)
			fmt.Printf(monthHeaderPrefixFormat, " ")
			fmt.Printf("%s", monthHeader)
			fmt.Printf(monthHeaderPrefixFormat, " ")
			fmt.Print("  ")
		}
		// Go to next line to print weekday names
		fmt.Println()

		// Print the weekday names
		for col := 1; col <= 3; col++ {
			// Print the Names
			for _, weekDay := range weekDays {
				fmt.Printf("%2s ", weekDay.dSName)
			}
			fmt.Print("  ")
		}
		// Go to next line to start printing day values
		fmt.Println()

		lastDayNo := [3]int{1, 1, 1}
		// Print the day values in 6 * 7 grid
		for mrow := 1; mrow <= 6; mrow++ {

			for mcol := 1; mcol <= 3; mcol++ {
				// Populate the month number in mNo and get the month struct
				mNo := mcol + (row-1)*3
				m := yearMonths[mNo-1]

				// Shift the date to correct position by using spaces in the first row
				if mrow == 1 {
					for j := 1; j <= m.startDayIndex; j++ {
						fmt.Printf("   ")
					}
				}

				// Print the day values after the shifted position and increment lastDayNo until
				// it becomes greater than month noOfDays. Whenever the dayPosition is greater
				// than or equal to seven 7 next row should be populated hence break the loop
				// Next iteration will start from the saved day in lastDayNo
				for lastDayNo[mcol-1] <= m.noOfDays {
					dayPosition := m.startDayIndex + lastDayNo[mcol-1]
					if dayPosition <= mrow*7 {
						if y.year == currYear && m.mNumber == currMonth && lastDayNo[mcol-1] == currDay {
							fmt.Printf("%s%2d%s ", invert, lastDayNo[mcol-1], reset)
						} else {
							fmt.Printf("%2d ", lastDayNo[mcol-1])
						}

						lastDayNo[mcol-1]++
					} else {
						break
					}
				}

				// Print spaces in the last row. If the last row does not contain any day then 7
				// spaces should be placed to offset the next month start. If the last row contains
				// a day then count of offsets should be calculated and put spaces for the count.
				monthLastDayIndex := m.noOfDays + m.startDayIndex
				if monthLastDayIndex/7 < (mrow-1) || (monthLastDayIndex/7 == (mrow-1) && monthLastDayIndex%7 == 0) {
					for i := 1; i <= 7; i++ {
						fmt.Print("   ")
					}
				} else if lastDayNo[mcol-1] > m.noOfDays {
					lastPosition := m.startDayIndex + lastDayNo[mcol-1] - 1
					shiftOffset := 0
					if lastPosition%7 != 0 {
						shiftOffset = lastPosition % 7
					} else {
						shiftOffset = 7
					}
					for i := 1; i <= (7 - shiftOffset); i++ {
						fmt.Print("   ")
					}
				}

				fmt.Print("  ")
			}

			fmt.Println()
		}

		fmt.Println()
	}

	// Send terminal prompt to next line
	fmt.Println()
}
