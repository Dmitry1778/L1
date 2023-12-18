package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan [5]int) // создаем первый канал массива с размером 5
	ch2 := make(chan [5]int) //создаем второй канал массива с размером 5
	x := [5]int{1, 2, 3, 4, 5}
	// горутина один
	go func() {
		ch1 <- x
	}()
	fmt.Println(<-ch1)
	// горутина два
	go func() {
		for i := 0; i < len(x); i++ {
			x[i] *= 2
		}
		ch2 <- x
	}()
	fmt.Println(<-ch2)
}
