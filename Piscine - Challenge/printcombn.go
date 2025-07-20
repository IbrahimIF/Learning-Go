package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n > 10 {
		return
	}

	var recursive func(depth int, start int, combo []rune)

	recursive = func(depth int, start int, combo []rune) {
		if depth == n {
			for _, r := range combo {
				z01.PrintRune(r)
			}

			if combo[0] != rune('0'+(10-n)) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}

		for i := start; i <= 9; i++ {
			recursive(depth+1, i+1, append(combo, rune('0'+i)))
		}
	}

	recursive(0, 0, []rune{})
	z01.PrintRune('\n')
}
