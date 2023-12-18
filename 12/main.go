package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(removeDuplicate(str))
	fmt.Println(removeDup(str))
}

func removeDuplicate(array []string) []string {
	var result []string
	keys := make(map[string]bool)
	for _, entry := range array {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}

func removeDup(arr []string) []string {
	m := make(map[string]bool)

	for _, x := range arr {
		m[x] = true
	}
	var res []string
	for x, _ := range m {
		res = append(res, x)
	}
	return res
}
