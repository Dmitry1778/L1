package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{1, 3, 4, 2, 6, 12, 9, 13, 10, 7, 8}
	target := 3
	current := len(slice) / 2 //(12) число с которым будем работать
	slices.Sort(slice)
	if len(slice) == 0 {
		return
	}
	fmt.Println(slice)
	fmt.Println(binSearch(slice, target, current)) // передали массив сортированный, число которое ищем(1), текущее число(12)
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
