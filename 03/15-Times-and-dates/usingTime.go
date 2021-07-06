package main

import (
	"fmt"
	"time"
)

func main() {

	// time.Now().Unix() returns the UNIX epoch time, which is the number of seconds
	// that have elapsed since 00:00:00 UTC, 1 January, 1970
	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()

	// The Format functioin allows us to convert a time variable to another format
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()

	// Sub() allows us to find the time difference between two times.
	fmt.Println("Time difference:", t1.Sub(t))

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)

	loc, err := time.LoadLocation("Europe/Athens")
	if err != nil {
		fmt.Println("Cannot find location", err)
	}

	athensTime := t.In(loc)
	fmt.Println("Athens:", athensTime)

}
