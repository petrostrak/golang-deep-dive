package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// This code will tell us how to read and store unstructured JSON data. The
// critical thing to remember is that unstructured JSON data is put into
// Go maps instead of Go structures.
func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename")
		return
	}

	filename := arguments[1]
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("cannot read file", err)
		return
	}

	var parsedData map[string]interface{}
	if err := json.Unmarshal(b, &parsedData); err != nil {
		fmt.Println("cannot unmarshal to parsedData", err)
		return
	}

	for key, value := range parsedData {
		fmt.Println("key:", key, "value:", value)
	}
}
