package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func Solution(s string) string {
	var result []rune
	var prev rune
	isPrevDigit := false

	for i, r := range s {
		if unicode.IsDigit(r) {
			if isPrevDigit || i == 0 {
				return "Некорректная строка"
			}
			count, _ := strconv.Atoi(string(r))
			for j := 0; j < count-1; j++ { 
				result = append(result, prev)
			}
			isPrevDigit = true
		} else {
			result = append(result, r)
			prev = r
			isPrevDigit = false
		}
	}

	return string(result)
}

func main() {
	fmt.Println(Solution("a4bc2d5e"))
	fmt.Println(Solution("abcd"))
	fmt.Println(Solution("45"))
	fmt.Println(Solution(""))
}