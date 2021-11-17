package analyzer

import (
	"fmt"
	"goshallowdive/helper"
	"sort"
	"strings"
)

func RemoveUnwantedChars(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}

	return result.String()
}

func CountWords(s string) map[string]int {
	RemoveUnwantedChars(s)
	words := strings.Fields(s)
	occurences := make(map[string]int, 2)

	for i := 0; i < len(words); i++ {
		elem, ok := occurences[words[i]]

		if ok == false {
			occurences[words[i]] = 1
		} else {
			occurences[words[i]] = elem + 1
		}
	}

	return occurences
}

func TakeMostCommonWords(text string, size int) []helper.Pair {
	fmt.Println("Started file analysis")

	wordsMap := CountWords(text)

	var ss []helper.Pair
	for k, v := range wordsMap {
		ss = append(ss, helper.Pair{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	cutSlice := ss[0:size]
	return cutSlice
}
