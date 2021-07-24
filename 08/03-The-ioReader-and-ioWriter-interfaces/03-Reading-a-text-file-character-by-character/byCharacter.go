package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func byCharacter(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("cannot read file %s", err)
			break
		}

		for _, x := range line {
			fmt.Print(string(x))
		}
	}

	return nil
}

func main() {

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byCharacter <file1> [<file2>...\n]")
		return
	}

	for _, file := range flag.Args() {
		if err := byCharacter(file); err != nil {
			fmt.Println(err)
		}
	}

}
