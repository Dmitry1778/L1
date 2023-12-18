package main

import "fmt"

func main() {
	array := []int{3, 1, 2, 5, 4}
	//str := []string{"1", "2", "4", "3", "5", "7", "10", "9", "11", "12", "8"}
	fmt.Println(quickSort(array, 0, len(array)-1))

	//sort.Ints(arr)
	//fmt.Println("ints", arr)
	//
	//slices.Sort(arr)
	//fmt.Println("slices", arr)
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		var p int
		arr, p = partition(arr, left, right)
		arr = quickSort(arr, left, p-1)
		arr = quickSort(arr, p+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) ([]int, int) {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}
