package main

import (
	"flag"
	"fmt"
	"strings"
)

type NameFlag struct {
	Names []string
}

func (s *NameFlag) GetNames() []string {
	return s.Names
}

func (s *NameFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NameFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names flag more than onces")
	}

	names := strings.Split(v, ",")
	s.Names = append(s.Names, names...)

	return nil
}

func main() {

	var namyNames NameFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Petros", "The name")
	flag.Var(&namyNames, "names", "Comma-separated list")

	flag.Parse()
	fmt.Println("-k", *minusK)
	fmt.Println("-o", *minusO)

	for i, item := range namyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining command line arguments")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}

}
