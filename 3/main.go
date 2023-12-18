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

// Просто метод для реализации суммы квадратов БЕЗ конкурентных вычислений.
func valueSquare(arr []int) int {
	var summ int
	for i := 0; i < len(arr); i++ {
		summ += arr[i] * arr[i]
	}
	return summ
}

func channelSquare(arr []int) int {
	var squares int
	ch := make(chan int) // канал типа int
	// Горутина которая пробегается по срезу и записывает данные в канал умножая число само на себя получая квадрат числа
	go func() {
		for _, v := range arr {
			ch <- v * v
		}
		close(ch)
	}()
	// цикл по каналу ch где в squares записываем сумму квадратов в переменную squares
	for v := range ch {
		squares += v
	}
	return squares
}

func waitSummSquare(array []int) int {
	var wg sync.WaitGroup // механизм ожидания завершения группы задач
	var summ int
	array = []int{2, 4, 6, 8, 10}
	for i := 0; i < len(array); i++ {
		wg.Add(1) // добавляем указанное количество горутин в счетчик ожидания.
		// горутина с анонимной фукцией
		go func(i int) {
			defer wg.Done() // при каждм вызове метода счетчик синхронизации уменьшается на 1.
			summ += array[i] * array[i]
		}(i) // передача значения переменной i в функцию func
	}
	wg.Wait() // блокируется до тех пор, пока счетчик синхронизации не станет равным нулю.
	return summ
}
