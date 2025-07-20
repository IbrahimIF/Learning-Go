package piscine

func NRune(s string, n int) rune {
	in := n - 1
	if in < len(s) && in > -1 {
		return []rune(s)[in]
	}
	return 0
}
