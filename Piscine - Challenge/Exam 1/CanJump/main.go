package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println("--- Testing CanJump ---")

	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Printf("CanJump(%v): %v\n", input1, piscine.CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Printf("CanJump(%v): %v\n", input2, piscine.CanJump(input2))

	input3 := []uint{0}
	fmt.Printf("CanJump(%v): %v\n", input3, piscine.CanJump(input3))

	input4 := []uint{}
	fmt.Printf("CanJump(%v): %v\n", input4, piscine.CanJump(input4))

	input5 := []uint{1}
	fmt.Printf("CanJump(%v): %v\n", input5, piscine.CanJump(input5))

	input6 := []uint{0, 1}
	fmt.Printf("CanJump(%v): %v\n", input6, piscine.CanJump(input6))

	input7 := []uint{1, 0}
	fmt.Printf("CanJump(%v): %v\n", input7, piscine.CanJump(input7))

	input8 := []uint{2, 0, 0, 1}
	fmt.Printf("CanJump(%v): %v\n", input8, piscine.CanJump(input8))


	fmt.Println("\n--- Testing FirstWord ---")

	fmt.Printf("FirstWord(\"Hello World\"): \"%s\"\n", piscine.FirstWord("Hello World"))        
	fmt.Printf("FirstWord(\"  Go  Programming\"): \"%s\"\n", piscine.FirstWord("  Go  Programming"))
	fmt.Printf("FirstWord(\"SingleWord\"): \"%s\"\n", piscine.FirstWord("SingleWord"))
	fmt.Printf("FirstWord(\"\"): \"%s\"\n", piscine.FirstWord(""))
	fmt.Printf("FirstWord(\"   \"): \"%s\"\n", piscine.FirstWord("   "))
	fmt.Printf("FirstWord(\"  Leading spaces\"): \"%s\"\n", piscine.FirstWord("  Leading spaces")) 
	fmt.Printf("FirstWord(\"Trailing spaces   \"): \"%s\"\n", piscine.FirstWord("Trailing spaces   ")) 
	fmt.Printf("FirstWord(\"One\tTwo\"): \"%s\"\n", piscine.FirstWord("One\tTwo")) 
	fmt.Printf("FirstWord(\"Hello,World!\"): \"%s\"\n", piscine.FirstWord("Hello,World!")) 

}