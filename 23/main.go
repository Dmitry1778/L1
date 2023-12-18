package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	i := 2 // выбираем i-элемент
	fmt.Println(removeElement(slice, i))
}

func removeElement(arr []int, i int) []int {
	var result []int
	fmt.Println(arr[i:])
	fmt.Println(arr[i+1:])
	copy(arr[i:], arr[i+1:]) //эта функция копирует из элементы из исходного в целевой фрагмент сдвигает элементы влево удаляя элемент с индексом i
	result = arr[:len(arr)-1]
	return result
}
