package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	// TODO: Fill in definition
	Id   string `json:"id"`
	Name string `json:"name`
}

// imporDtata reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {

	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatal("error occurred while reading data from json")
	}

	var data []raffleEntry
	e := json.Unmarshal(file, &data)
	if e != nil {
		log.Fatal("error occurred while converting from data to json")
	}

	return data
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
