package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	var compactedSlice []string
	count := 0
	for _, value := range s {
		if value != "" {
			compactedSlice = append(compactedSlice, value)
		}
	}
	*ptr = compactedSlice
	return count
}
