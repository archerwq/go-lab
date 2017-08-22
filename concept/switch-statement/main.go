package main

import (
	"fmt"
	"runtime"
	"time"
)

func osType() string {
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X."
	case "linux":
		return "Linux."
	default:
		return os + "."
	}
}

func whenIsSat() string {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		return "Today."
	case today + 1:
		return "Tomorrow."
	case today + 2:
		return "In two days."
	default:
		return "Too far away."
	}
}

func greeting() string {
	t := time.Now()
	// Switch without a condition, a clean way to write long if-then-else chains.
	switch {
	case t.Hour() < 12:
		return "Good morning!"
	case t.Hour() < 17:
		return "Good afternoon!"
	default:
		return "Good evening!"
	}
}

func main() {
	fmt.Println(osType())
	fmt.Println(whenIsSat())
	fmt.Println(greeting())
}
