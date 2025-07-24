package main

import (
	"os"

	"github.com/01-edu/z01"
	"strconv"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	if num <= 0 {
		printStr("false")
		z01.PrintRune('\n')
		return
	}

	isPowerOfTwo := false

	for i := 1; i <= num; i *= 2 {
		if i == num {
			isPowerOfTwo = true
			break
		}
	}
	if isPowerOfTwo {
		printStr("true")
	} else {
		printStr("false")
	}
	z01.PrintRune('\n')
}
