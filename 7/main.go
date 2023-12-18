package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

type SafeNumbers struct {
	sync.RWMutex
	numbers map[int]int
}

func (s *SafeNumbers) Add(num int) {
	s.Lock()
	defer s.Unlock()
	s.numbers[num] = num
}

func (s *SafeNumbers) Get(num int) (int, error) {
	s.RLock()
	defer s.RUnlock()
	if number, ok := s.numbers[num]; ok {
		return number, nil
	}
	return 0, errors.New("Number does not exists")
}

func generateNumbersMap() {
	var wg sync.WaitGroup
	safeNumbers := &SafeNumbers{
		numbers: map[int]int{},
	}
	//write
	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			safeNumbers.Add(i)
			wg.Done()
		}
	}()
	//read
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

//https://webdevstation.com/posts/concurrent-map-writing-and-reading-in-go/
