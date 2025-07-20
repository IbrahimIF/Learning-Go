package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	count := max - min
	arr := make([]int, count)
	j := 0
	for i := min; i < max; i++ {
		arr[j] = i
		j++
	}
	return arr
}
