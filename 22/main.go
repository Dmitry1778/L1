package main

import (
	"fmt"
)

func main() {
	var a = 1048580
	var b = 5012
	fmt.Println(a, "*", b, "=", mul(a, b))
	fmt.Println(a, "/", b, "=", div(a, b))
	fmt.Println(a, "+", b, "=", fold(a, b))
	fmt.Println(a, "-", b, "=", sub(a, b))
}

// a*b
func mul(a, b int) int {
	return a * b
}

// a/b
func div(a, b int) int {
	return a / b
}

// a+b
func fold(a, b int) int {
	return a + b
}

// a-b
func sub(a, b int) int {
	return a - b
}
