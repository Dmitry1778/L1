package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	array := [5]int{2, 4, 6, 8, 10}

	summ := waitSummSquare(array[:])
	fmt.Print("wait group result:", summ, " time:")
	time.Sleep(time.Millisecond)
	fmt.Println(time.Now().Sub(start).Seconds())

	// value := valueSquare(array[:])
	//fmt.Println(value)
	//time.Sleep(time.Millisecond)
	//fmt.Println(time.Now().Sub(start).Seconds())

	res := channelSquare(array[:])
	fmt.Print("channel result:", res, " time:")
	time.Sleep(time.Millisecond)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func valueSquare(arr []int) int {
	var summ int
	for i := 0; i < len(arr); i++ {
		summ += arr[i] * arr[i]
	}
	return summ
}

func channelSquare(arr []int) int {
	var squares int
	ch := make(chan int)
	go func() {
		for _, v := range arr {
			ch <- v * v
		}
		close(ch)
	}()

	for v := range ch {
		squares += v
	}
	return squares
}

func waitSummSquare(array []int) int {
	var wg sync.WaitGroup
	var summ int
	array = []int{2, 4, 6, 8, 10}
	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			summ += array[i] * array[i]
		}(i)
	}
	wg.Wait()
	return summ
}
