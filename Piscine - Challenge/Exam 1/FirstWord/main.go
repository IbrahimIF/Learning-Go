package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FirstWord("Hello world"))
	fmt.Println(piscine.FirstWord("  leading spaces"))
	fmt.Println(piscine.FirstWord("onlyOneWord"))
}