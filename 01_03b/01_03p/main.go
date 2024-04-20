package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into

type raffleEntry struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// imporDtata reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	file, err := os.ReadFile(path)

	jsonData := []byte(`{"id":"11", "name":30}`)

	var d []raffleEntry
	e1 := json.Unmarshal(jsonData, &d)
	if e1 != nil {
		fmt.Println(" error occurred while converting the json data")
	}

	fmt.Println(d)
	if err != nil {
		log.Fatal("error occurred while reading the data from the file")
	}

	var data []raffleEntry

	e := json.Unmarshal(file, &data)
	if e != nil {
		log.Fatal(" error occurred while converting from text to json")
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
