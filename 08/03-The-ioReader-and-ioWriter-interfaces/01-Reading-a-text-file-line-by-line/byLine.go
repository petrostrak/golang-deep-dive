package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	err error
)

func lineByline(file string) error {
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
			fmt.Printf("error reading file %s\n", err)
			break
		}
		fmt.Println(line)
	}

	return nil
}

func main() {

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		if err := lineByline(file); err != nil {
			fmt.Println(err)
		}
	}

}
