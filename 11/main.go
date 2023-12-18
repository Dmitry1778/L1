package main

import "fmt"

func main() {
	var sets1 = []int{3, 2, 4, 7, 2, 3}
	var sets2 = []int{5, 2, 6, 7, 4, 5}

	array := findInters(sets1, sets2)

	result := removeDuplicate(array)

	fmt.Println("Пересечение:", result)
}

func findInters(arrOne []int, arrTwo []int) []int {
	var result []int
	for i := 0; i < len(arrOne); i++ {
		for j := 0; j < len(arrTwo); j++ {
			if arrOne[i] == arrTwo[j] {
				result = append(result, arrTwo[j])
			}
		}
	}
	return result
}

func removeDuplicate(array []int) []int {
	var result []int
	keys := make(map[int]bool)
	for _, entry := range array {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}
