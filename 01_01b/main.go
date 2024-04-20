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
	dt := target.Sub(time.Now()) / 24
	dt1 := time.Until(target).Hours() / 24
	fmt.Println(" the until function returns", dt1)
	fmt.Println(" the sub function returns", dt)

	return dt1
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
