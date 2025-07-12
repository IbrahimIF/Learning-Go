package piscine

func StrRev(s string) string {
	str := []byte(s)
	for i := 0; i <= len(s)-1; i++ {
		str[len(s)-1-i] = s[i]
	}
	return string(str)
}
