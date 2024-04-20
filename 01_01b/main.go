package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	dt, err := time.Parse(expectedFormat, target)
	if err != nil || time.Now().After(dt) {
		log.Fatal("date and time is invaid")
	}
	return dt

}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	fmt.Println("current time ", time.Now(), "target is :", target)
	dt := target.Sub(time.Now())
	fmt.Println(" the duration :", dt)
	return dt.Hours() / 24
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
