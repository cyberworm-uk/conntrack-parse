package main

import (
	"conntrack"
	"fmt"
	"os"
)

func main() {
	var err error
	var table []byte
	table, err = os.ReadFile("/dev/stdin")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, row := range conntrack.Parse(string(table)) {
		fmt.Printf("%s\n", row)
	}
}
