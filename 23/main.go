package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	i := 3
	fmt.Println(removeElement(slice, i))
}

func removeElement(arr []int, i int) []int {
	var result []int
	copy(arr[i:], arr[i+1:])
	result = arr[:len(arr)-1]
	return result
}
