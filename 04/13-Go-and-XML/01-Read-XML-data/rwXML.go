package main

import (
	"encoding/json"
	"encoding/xml"
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

func loadFromJSON(filename string, key interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(file).Decode(key); err != nil {
		return err
	}
	file.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]

	var myRecord Record
	if err := loadFromJSON(filename, &myRecord); err != nil {
		fmt.Println("cannot decode json file", err)
	}

	myRecord.Name = "Pit"
	fmt.Println(myRecord)

	xmlData, _ := xml.MarshalIndent(myRecord, "", "	")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("\nxmlData:", string(xmlData))

	data := &Record{}
	if err := xml.Unmarshal(xmlData, data); err != nil {
		fmt.Println("cannot unmarshal from XML", err)
		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("cannot marshal to JSON", err)
		return
	}

	_ = json.Unmarshal([]byte(b), &myRecord)
	fmt.Println("\nJSON:", myRecord)

}
