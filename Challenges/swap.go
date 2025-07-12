package piscine

func Swap(a *int, b *int) {
	ai := *a
	bi := *b
	*a = bi
	*b = ai
}
