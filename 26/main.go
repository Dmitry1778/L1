package main

import "fmt"

func main() {
	fmt.Println(checkingUniq("abcd"))
	fmt.Println(checkingUniq("abCdefAaf"))
	fmt.Println(checkingUniq("aabcd"))
}

func checkingUniq(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
