package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "snow dog sun"
	reverse := reverseWord(s)
	fmt.Println(reverse)
}

func reverseWord(str string) string {
	var result []string
	words := strings.Fields(str)
	for j := len(words) - 1; j >= 0; j-- {
		result = append(result, words[j])
	}
	return strings.Join(result, " ")
}
