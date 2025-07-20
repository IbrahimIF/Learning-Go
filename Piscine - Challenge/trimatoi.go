package piscine

func TrimAtoi(s string) int {
	sign := 1
	foundSign := false
	numStarted := false
	num := 0

	for _, r := range s {
		if r == '-' && !numStarted && !foundSign {
			sign = -1
			foundSign = true
		} else if r >= '0' && r <= '9' {
			numStarted = true
			num = num*10 + int(r-'0')
		} else if (r == '+' || r == '-') && numStarted {
			continue
		}
	}

	if !numStarted {
		return 0
	}
	return num * sign
}
