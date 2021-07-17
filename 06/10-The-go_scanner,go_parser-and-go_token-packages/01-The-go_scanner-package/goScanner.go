package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments")
		return
	}

	// the source file that is going to be tokenized is stored in the file variable, whereas its
	// contents are stored in f.
	for _, file := range arguments[1:] {
		f, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		One := token.NewFileSet()
		files := One.AddFile(file, One.Base(), len(f))

		var myScanner scanner.Scanner
		myScanner.Init(files, f, nil, scanner.ScanComments)

		for {
			pos, tok, lit := myScanner.Scan()
			if tok == token.EOF {
				break
			}
			fmt.Printf("%s\t%s\t%q\n", One.Position(pos), tok, lit)
		}

	}

}
