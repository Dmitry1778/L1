package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{1, 3, 4, 2, 6, 12, 9, 13, 10, 7, 8}
	target := 1
	current := len(slice) / 2
	slices.Sort(slice)
	if len(slice) == 0 {
		return
	}
	fmt.Println(binSearch(slice, target, current))
}

func binSearch(slice []int, target, current int) int {
	var result int
	for result != target {
		if target < slice[current] {
			result = current
			current = slice[current] / 2
		}
	}
	return result
}

//https://www.youtube.com/watch?v=uHdS2LA-yn8
