package main

import (
	"encoding/json"
	"fmt"
)

type Record struct {
	Name, Surname string
	Tel           []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
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

	b, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println("cannot marshal to JSON", err)
		return
	}

	fmt.Println(string(b))

	var myRec Record
	if err := json.Unmarshal(b, &myRec); err != nil {
		fmt.Println("cannot unmarshal JSON", err)
		return
	}

	fmt.Println(myRec)

}
