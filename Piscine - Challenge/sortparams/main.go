package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := os.Args[1:]

	for i := 0; i <= len(str)-1; i++ {
		for j := 0; j < len(str)-i-1; j++ {
			if str[j] > str[j+1] {
				str[j], str[j+1] = str[j+1], str[j]
			}
		}
	}

	for _, word := range str {
		for _, char := range word {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
