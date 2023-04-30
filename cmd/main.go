package main

import (
	"gocalendar"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Usage descrption and information
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Simple Command Line Calendar Tool Utility. Inspired by cal tool in Linux.")
		fmt.Fprintln(flag.CommandLine.Output(), "Author: Suman Das")
		fmt.Fprintln(flag.CommandLine.Output(), "Year: 2022")
		fmt.Fprintf(flag.CommandLine.Output(), "\nUsage; %s \n", os.Args[0])
		flag.PrintDefaults()
	}

	// Get Current Year based on time package
	now := time.Now()
	currYear := now.Year()
	currMonth := int(now.Month())

	// Command line argument flag -y
	year := flag.Int("y", 0, "Calendar Year Input")
	month := flag.Int("m", 0, "Calendar Month Input")

	flag.Parse()

	switch {
	case *year != 0 && *month != 0:
		m, err := calendar.NewMonth(*year, *month)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Show the Month Calendar
		m.Print(now)
	case *year != 0 && *month == 0:
		y, err := calendar.NewYear(*year)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Show the Year Calendar
		y.Print(now)
	default:
		m, err := calendar.NewMonth(currYear, currMonth)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Show the Month Calendar
		m.Print(now)
	}

}
