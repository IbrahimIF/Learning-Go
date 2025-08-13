package piscine

import (
	"sort"
	"strings"
)

func MaxWordCountN(text string, n int) map[string]int {
	rawWords := strings.Split(text, " ")
	words := []string{}
	for _, word := range rawWords {
		if word != "" {
			words = append(words, word)
		}
	}

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	type pair struct {
		word  string
		count int
	}
	var pairs []pair
	for k, v := range freq {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].word < pairs[j].word
		}
		return pairs[i].count > pairs[j].count
	})

	result := make(map[string]int)
	for i := 0; i < n && i < len(pairs); i++ {
		result[pairs[i].word] = pairs[i].count
	}

	return result
}