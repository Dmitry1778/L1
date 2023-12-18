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
	words := strings.Fields(str) // разбивает строку «str» на фрагмент строк, где каждый элемент фрагмента представляет слово из исходной строки.
	for j := len(words) - 1; j >= 0; j-- {
		result = append(result, words[j]) // последнее слово на первое место
	}
	return strings.Join(result, " ") // объединяет элементы фрагмента строки `result` в одну строку, где каждый элемент отделяется пробелом (" ").
}
