package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	var digits []byte
	for n > 0 {
		d := byte(n % 10)
		digits = append(digits, '0'+d)
		n /= 10
	}
	if isNegative {
		digits = append(digits, '-')
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}
