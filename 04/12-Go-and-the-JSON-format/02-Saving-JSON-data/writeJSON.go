package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name, Surname string
	Tel           []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

var (
	FILENAME = "/home/petros_trak/github.com/petrostrak/golang-deep-dive/04/12-Go-and-the-JSON-format/02-Saving-JSON-data/myRecord.json"
)

func saveToJSON(filename *os.File, key interface{}) {
	if err := json.NewEncoder(filename).Encode(key); err != nil {
		fmt.Println("cannot encode to JSON format", err)
		return
	}
}

func main() {
	myRecord := Record{
		Name:    "Petros",
		Surname: "Trak",
		Tel: []Telephone{
			{Mobile: true, Number: "123-4566"},
			{Mobile: true, Number: "555-1234"},
			{Mobile: false, Number: "abc-1234"},
		},
	}

	f, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cannot open file to save JSON data", err)
	}
	defer f.Close()

	saveToJSON(f, myRecord)
}
