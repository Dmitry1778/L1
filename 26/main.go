package main

import "fmt"

func main() {
	fmt.Println(checkingUniq("abcd"))
	fmt.Println(checkingUniq("abCdefAaf"))
	fmt.Println(checkingUniq("aabcd"))
}

func checkingUniq(s string) bool {
	// внешний цикл для первой буквы чтобы потом сравнить с другими
	for i := 0; i < len(s); i++ {
		// внутренний цикл для всех остальных символов чтобы сравить с первым символом(главным)
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
