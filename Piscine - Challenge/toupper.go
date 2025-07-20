package piscine

func ToUpper(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			result += string(rune(s[i]) - 32)
			continue
		}
		result += string(s[i])
	}
	return result
}
