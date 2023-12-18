package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	array := [5]int{2, 4, 6, 8, 10}
	sq := channelSquare(array[:])
	fmt.Println(sq)
	time.Sleep(time.Millisecond)
	fmt.Println(time.Now().Sub(start).Seconds())

	//waitValueSquare(array[:])
}

func channelSquare(arr []int) []int {
	squares := make([]int, 0, len(arr)) // Создали срез с длинной 0 и ёмкостью равной длине среза "arr"
	ch := make(chan int)                // Канал типа int
	// Горутина которая пробегается по срезу и записывает данные в канал умножая число само на себя получая квадрат числа
	go func() {
		for _, v := range arr {
			ch <- v * v // запись в канал
		}
		close(ch) // Закрываем канал, т.к. данных больше не будет
	}()
	// Цикл по каналу
	for v := range ch {
		squares = append(squares, v) // добавлаем данные в массив
	}
	return squares
}

func waitValueSquare(array []int) {
	start := time.Now()
	var wg sync.WaitGroup // механизм ожидания завершения группы задач
	//wg.Add(5)
	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(array[i] * array[i])
			time.Sleep(time.Nanosecond)
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start).Seconds())
	fmt.Println("exit")
}

//https://www.youtube.com/watch?v=RaMQrV7x-1k&list=WL&index=2&t=431s - (Изучаем Golang. Урок №20. Concurrency (2). WaitGroup. Data Race. Muxtex/RWMutex)
//https://www.youtube.com/watch?v=wHQBMDInWEg&t=467s про горутины расскажет этот (Maxim)
//https://www.youtube.com/watch?v=aJ74a1gydbQ&list=WL&index=162&t=684s
