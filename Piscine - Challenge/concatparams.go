package piscine

func ConcatParams(args []string) string {
	str := ""
	for i := range args {
		if i+1 == len(args) {
			str = str + args[i]
		} else {
			str = str + args[i] + "\n"
		}
	}
	return str
}
