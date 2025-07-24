package piscine

func Rot14(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		r := s[i]
		if s[i] >= 'a' && s[i] <= 'z' {
			if s[i]+14 > 122 {
				r = ((s[i] + 14) - 122) + 96
			} else {
				r = s[i] + 14
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			if s[i]+14 > 90 {
				r = ((s[i] + 14) - 90) + 64
			} else {
				r = s[i] + 14
			}
		} else {
			r = s[i]
		}

		result += string(r)
	}
	return result
}
