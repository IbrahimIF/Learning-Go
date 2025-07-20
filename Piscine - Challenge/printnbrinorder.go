package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var counts [10]int
	for n > 0 {
		digit := n % 10
		counts[digit]++
		n = n / 10
	}

	for digit := 0; digit <= 9; digit++ {
		for i := 0; i < counts[digit]; i++ {
			z01.PrintRune(rune(digit + '0'))
		}
	}
}
