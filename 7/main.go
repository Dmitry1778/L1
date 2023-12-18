package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

type SafeNumbers struct {
	sync.RWMutex             // механизм для получения исключительной блокировки нужен чтобы защитить данные от конфликта и потери данных
	numbers      map[int]int // переменная которую нужно защитить во время работы с ней несколькими горутинами
}

func (s *SafeNumbers) Add(num int) {
	// благоря Lock мы безопасно записываем переменную после чего блокировка снимается
	s.Lock()
	defer s.Unlock()
	s.numbers[num] = num
}

func (s *SafeNumbers) Get(num int) (int, error) {
	//RLock - неисклюзивная блокировка так как оно не изменяется то все горутины могут читать значение в данном случае это горутина 2
	s.RLock()
	defer s.RUnlock()
	if number, ok := s.numbers[num]; ok {
		return number, nil
	}
	return 0, errors.New("Number does not exists")
}

func generateNumbersMap() {
	var wg sync.WaitGroup // механизм ожидания завершения группы задач
	safeNumbers := &SafeNumbers{
		numbers: map[int]int{},
	}
	//write
	// первая горутина записывает данные в карту
	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			safeNumbers.Add(i)
			wg.Done()
		}
	}()
	//read
	// вторая горутина читает данные из карты
	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			number, err := safeNumbers.Get(i)
			if err != nil {
				log.Print(err)
			} else {
				log.Print(number)
			}
			wg.Done()
		}
	}()
	wg.Wait()
}

//func generateNumbersMap() {
//	wg := sync.WaitGroup{}
//	// Init our "safe" numbers map struct.
//	safeNumbers := &SafeNumbers{
//		numbers: map[int]int{},
//	}
//	// Write.
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			safeNumbers.Add(i)
//		}(i)
//	}
//	// Read.
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			number, err := safeNumbers.Get(i)
//			if err != nil {
//				log.Print(err)
//			} else {
//				log.Print(number)
//			}
//		}(i)
//	}
//
//	wg.Wait()
//}

func main() {
	generateNumbersMap()
	time.Sleep(time.Millisecond * 500)
}

//2,5,6
//https://webdevstation.com/posts/concurrent-map-writing-and-reading-in-go/
