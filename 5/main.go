package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	const N = 5                          // время по которому программа должна завершаться
	ch := make(chan string, 1)           // буфферизированный канал
	timer := time.After(time.Second * N) // сама функция которая создает таймер, который будет срабатывать через N секунд.

	// чтение из канала
	read(ch)

	// цикл для отправления значений в канал
	for {
		select {
		case <-timer:
			close(ch)
			fmt.Println("time's out", time.Now().Sub(start).Seconds(), "second")
			time.Sleep(time.Millisecond * 100)
			return
		default:
			ch <- "Hello World!"
		}
		time.Sleep(time.Second)
	}
}

func read(ch chan string) {
	go func() {
		for writeChannel := range ch {
			fmt.Println("output:", writeChannel)
		}
	}()
}

//https://www.youtube.com/watch?v=kSfR5LWKUrM&t=900s
