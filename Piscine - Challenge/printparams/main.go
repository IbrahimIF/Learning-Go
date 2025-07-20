package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := os.Args[1:]

	for _, word := range str {
		for _, char := range word {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
