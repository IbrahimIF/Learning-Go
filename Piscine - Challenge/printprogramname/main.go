package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]
	s := ""
	for _, curr := range programName {
		s += string(curr)
	}

	for _, curr := range programName[2:] {
		z01.PrintRune(curr)
	}
	z01.PrintRune('\n')
}
