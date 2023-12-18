package main

import "fmt"

func main() {
	num := int64(48)

	position := 3 // позиция в которую хотим установить 1 или 0

	bit := 1 // 1 или 0

	fmt.Printf("%08b\n", uint64(num)) // первоначальный вид

	result := setBit(num, bit, position)
	fmt.Println(result)
	fmt.Printf("%08b\n", uint64(result))
}

func setBit(num int64, bitValue, pos int) int64 {
	// Если позиция в которую мы хотим установить значение будет false то используем (объединение или set)
	if setOrUnset(bitValue) {
		return num &^ (1 << pos) // и не unset
	}
	return num | (1 << pos) // объединение или set
}

func setOrUnset(b int) bool {
	return b == 0
}
