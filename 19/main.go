package main

import (
	"fmt"
)

func main() {
	var s = "главрыба"
	fmt.Println(reverseString(s))
	fmt.Println(reverseStr(s))
}

// Однако, когда мы хотим манипулировать отдельными символами внутри строки, мы используем руны.
// Руны представляют собой кодовые точки Юникода, которые позволяют нам обрабатывать символы разных языков и сценариев.
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// толкаем вправо каждый элемент
func reverseStr(s string) string {
	var str string
	for _, v := range s {
		str = string(v) + str
	}
	return str
}

// rune working https://www.youtube.com/watch?v=wUQB74nNxos
