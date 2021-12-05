package main

import (
	"fmt"
	"time"
)

func withTime(fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	fmt.Println("millis:", duration.Milliseconds())
}

func main() {
	// day1()
	// day2()
	// day3()
	// day4()
	withTime(day5)
}
