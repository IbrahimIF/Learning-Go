package piscine

func SplitWhiteSpaces(s string) []string {
	newstr := s + " "
	strarr := []string{}
	count := 0
	for i := 0; i < len(newstr); i++ {
		if string(newstr[i]) == " " {
			if count < i {
				strarr = append(strarr, string(newstr[count:i]))
			}
			count = i + 1
		}
	}
	return strarr
}
