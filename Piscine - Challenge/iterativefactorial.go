package piscine

func IterativeFactorial(nb int) int {
	calc := 1
	if nb == 0 {
		return 1
	}
	if nb < 0 || nb > 21 {
		return 0
	} else {
		for i := nb; i >= 1; i-- {
			calc = i * calc
		}
		return calc
	}
}
