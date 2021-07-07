package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

// An IPv4 address, or simply an IP address, has four discrete parts. As an IPv4 address is
// stored using 8-bit binary numbers, each part can have values from 0 (which is 00000000 in
// binary format) to 255 (which is equal to 11111111 in binary format)
//
// findIp contains the definition of the regex that will help us to discover an IPv4 address
// inside a function.
func findIP(input string) string {

	// Be aware of the fact that the decimal values of an IPv4 address cannot be larger than
	// 255. A valid IPv4 address can begin with 25 and end with 0, 1, 2, 3, 4 or 5 because that
	// is the biggest 8-bit binary number (25[0-5]), or it can begin with 2 followed by 0, 1, 2,
	// 3 or 4 and end with 0, 1, 2, 3, 4, 5, 6, 7, 8 or 9 (2[0-4][0-9]). Alternatively, it can
	// begin with 1 followed by two more digits from 0-9 (1[0-9][0-9]). The last digit, which is
	// optional, can be from 1 to 9 and the second, which is mandatory, can be from 0-9 ([1-9]?[0-9])
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				break
			}

			// for each line of the input file, we call the findIP function
			ip := findIP(line)

			// The net.ParseIP double checks that we are dealing with a valid
			// IPv4 address. If the net.ParseIP is successful, we print the
			// IPv4 address we just found.
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			fmt.Println(ip)
		}
	}

}
