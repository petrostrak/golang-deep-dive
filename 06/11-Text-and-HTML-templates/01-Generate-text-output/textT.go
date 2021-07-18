package main

import (
	"fmt"
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need the template file")
		return
	}

	tFile := arguments[1]
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}

	var Entries []Entry

	// Create a slice of structures from DATA variable
	for _, i := range DATA {
		if len(i) == 2 {
			temp := Entry{i[0], i[1]}
			Entries = append(Entries, temp)
		}
	}

	// template.Must() is used for making the required initializations. Its returned
	// data type is Template which is a structure that holds the representation of
	// a parsed template.
	//
	// template.ParseGlob() reads the external template file.
	t := template.Must(template.ParseGlob(tFile))

	// t.Execute() does all the work, which includes processing the data and printing the
	// output to the desired file, which in this case is os.Stdout
	t.Execute(os.Stdout, Entries)

}

// Calculation the squares of some integers
// {{ range . }}
// The square of {{ printf "%d" .Number }} is {{ printf "%d" .Square }}
// {{ end }}
//
// The range keyword allows you to iterate over the lines of the input, which is given
// as a slice of structure. Plain text is printed as such, wheras variables and dynamic
// text must begin with {{ and end with }}. The fields of the structure are accessed as
// .Number and .Square. The printf commad is used to format the final output.
// The {{ range }} command is ended with {{ end }}.
