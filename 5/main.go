
package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words *[]string) map[string][]string {
	anagrams := make(map[string][]string)
	unique := make(map[string]bool)

	for _, word := range *words {
		lower := strings.ToLower(word)
		sortedWord := SortStr(lower)
		anagrams[sortedWord] = append(anagrams[sortedWord], lower)
	}

	result := make(map[string][]string)
	for _, group := range anagrams {
		if len(group) > 1 {
			sort.Strings(group)
			firstWord := group[0]
			if !unique[firstWord] {
				result[firstWord] = group
				unique[firstWord] = true
			}
		}
	}

	return result
}

func SortStr(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramGroups := findAnagrams(&words)
	for key, group := range anagramGroups {
		fmt.Printf("%s: %v\n", key, group)
	}
}