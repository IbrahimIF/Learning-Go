package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := os.Args[1:]

	for i := len(str) - 1; i >= 0; i-- {
		for _, char := range str[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
