package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var deadline string
	flag.StringVar(&deadline, "deadline", "", "the deadline for the timer in the foramt '1994-01-01 15:00:00'")
	flag.Parse()

	t, err := time.Parse("2006-01-02 15:04:05", deadline)
	if err != nil {
		fmt.Println("invalid deadline", err)
		os.Exit(1)
	}

	for {
		// Calculate the duration until the deadline
		duration := time.Until(t)

		if duration < 0 {
			fmt.Print("\rdeadline has passed")
			break
		}

		// Extract the remaining hours, minuts and seconds from the duration
		hours := int(duration.Hours())
		minutes := int(duration.Minutes()) % 60
		seconds := int(duration.Seconds()) % 60

		// Display the time remaining
		fmt.Printf("\rtime remaining: %02d:%02d:%02d", hours, minutes, seconds)

		// Wait for one second before updating the display.
		time.Sleep(time.Second)
	}
}
