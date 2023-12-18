package main

import "fmt"

//страница 78

func main() {
	num := int64(48)

	position := 3

	bit := 1

	fmt.Printf("%08b\n", uint64(num))

	result := setBit(num, bit, position)
	fmt.Println(result)
	fmt.Printf("%08b\n", uint64(result))
}

func setBit(num int64, bitValue, pos int) int64 {

	if setOrUnset(bitValue) {
		return num &^ (1 << pos) // и не unset
	}
	return num | (1 << pos) // объединение или set
}

func setOrUnset(b int) bool {
	return b == 0
}
