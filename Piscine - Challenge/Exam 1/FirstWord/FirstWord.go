package piscine

func FirstWord(s string) string {
    var j int
    for i := 0; i < len(s); i++ {
        if (s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') {
            j = i
            for j < len(s) {
                if s[j] == ' ' {
                    break
                }
                j++
            }
            return s[i:j]
        }
    }
    return ""
}