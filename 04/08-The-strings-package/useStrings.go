package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var (
	f = fmt.Printf
)

// The strings standard Go package allows you to manipulate UTF-8 strings in Go.
func main() {

	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE!"))
	f("%s\n", s.Title("tHis will bE a tItlE!"))
	f("EqualFold: %v\n", s.EqualFold("Petros", "pEtRos"))
	f("EqualFold: %v\n", s.EqualFold("Petros", "PETrO"))

	f("Prefix: %v\n", s.HasPrefix("Petros", "Pe"))
	f("Prefix: %v\n", s.HasPrefix("Petros", "pe"))
	f("Suffix: %v\n", s.HasSuffix("Petros", "os"))
	f("Suffix: %v\n", s.HasSuffix("Petros", "OS"))

	f("Index: %v\n", s.Index("Petros", "tr"))
	f("Index: %v\n", s.Index("Petros", "Tr"))
	f("Count: %v\n", s.Count("Petros", "e"))
	f("Count: %v\n", s.Count("Petros", "E"))
	f("Repeat: %v\n", s.Repeat("au", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s\n", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	f("Compare: %v\n", s.Compare("Petros", "PETROS"))
	f("Compare: %v\n", s.Compare("Petros", "Petros"))
	f("Compare: %v\n", s.Compare("PeTRos", "PeTRos"))

	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))

	f("%s\n", s.Split("Petros Trak", ""))

	// The strings.Replace takes 4 parameters. The 1st one is the string we want to process. The
	// 2nd contains the string that if found, will be replaced by the 3rd parameter. The last para-
	// meter is the maximum number of replacements that are allowed to happen.
	f("%s\n", s.Replace("Petros Trak", "", "_", 1))
	f("%s\n", s.Replace("Petros Trak", "", "_", 4))
	f("%s\n", s.Replace("Petros Trak", "", "_", 2))

	lines := []string{"Petros", "Trak", "Software", "Engineer"}
	f("Join: %s\n", s.Join(lines, "+++"))

	f("SplitAfter: %s\n", s.SplitAfter("123++456++7", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	// This removes every rune that is NOT a letter
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}
