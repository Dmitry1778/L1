package main

import "fmt"

func main() {
	var array = []int{1, 2}
	fmt.Println(reverse(array[:]))
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
