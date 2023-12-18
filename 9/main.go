package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan [5]int)
	ch2 := make(chan [5]int)
	x := [5]int{1, 2, 3, 4, 5}

	go func() {
		ch1 <- x
	}()
	fmt.Println(<-ch1)
	go func() {
		for i := 0; i < len(x); i++ {
			x[i] *= 2
		}
		ch2 <- x
	}()
	fmt.Println(<-ch2)
}
