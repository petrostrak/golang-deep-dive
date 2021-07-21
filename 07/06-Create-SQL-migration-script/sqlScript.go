package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	arguments := os.Args
	if len(arguments) <= 2 {
		fmt.Println("The program needs 2 input files")
		os.Exit(1)
	}

	T_CREDITOR_EMAILS := arguments[1]
	T_CODE_CREDITOR_UUIDs := arguments[2]

	f1, err := os.Open(T_CREDITOR_EMAILS)
	if err != nil {
		fmt.Printf("cannot open file: %s, err: %v", T_CREDITOR_EMAILS, err)
	}
	defer f1.Close()

	f2, err := os.Open(T_CODE_CREDITOR_UUIDs)
	if err != nil {
		fmt.Printf("cannot open file: %s, err: %v", T_CODE_CREDITOR_UUIDs, err)
	}
	defer f2.Close()

	r1 := bufio.NewReader(f1)
	r2 := bufio.NewReader(f2)
	var sb strings.Builder
	for {
		line1, err1 := r1.ReadString('\n')
		line2, err2 := r2.ReadString('\n')
		if err1 == io.EOF || err2 == io.EOF {
			break
		} else if err != nil {
			fmt.Println("cannot read file")
		}
		fmt.Fprintf(&sb, "INTO T_CREDITOR_EMAILS(uuid, email, version, fk_creditor) VALUES('%s' ,'mapistest@eurodyn.com', '1', '%s')\n", strings.TrimSuffix(line1, "\n"), strings.TrimSuffix(line2, "\n"))
	}

	f3, err := os.Create("sqlScript")
	if err != nil {
		fmt.Println("could not create file")
	}
	defer f3.Close()

	w := bufio.NewWriter(f3)
	n, err := w.WriteString(sb.String())
	if err != nil {
		fmt.Println("could not write to file")
	}
	fmt.Printf("wrote %d bytes\n", n)
}
