package main

import "fmt"

// Пересече́ние мно́жеств в теории множеств — это множество, которому принадлежат те и только те элементы, которые одновременно принадлежат всем данным множествам.
func main() {
	var sets1 = []int{3, 2, 4, 7, 2, 3}
	var sets2 = []int{5, 2, 6, 7, 4, 5}
	// ищем пересечения
	array := findInters(sets1, sets2)
	// убираем дубликаты
	result := removeDuplicate(array)
	// выводим результат
	fmt.Println("Пересечение:", result)
}

func findInters(arrOne []int, arrTwo []int) []int {
	var result []int
	for i := 0; i < len(arrOne); i++ {
		for j := 0; j < len(arrTwo); j++ {
			if arrOne[i] == arrTwo[j] { // если они одинаковые то записываем их в результат (в массив)
				result = append(result, arrTwo[j])
			}
		}
	}
	return result
}

// так как у нас могут быть дубликаты при поиске пересечение убираем их
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
