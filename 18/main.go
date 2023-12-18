package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	num int
}

func main() {
	c := newInit()
	fmt.Println(c.init())

}

func newInit() *counter {
	return &counter{}
}

// Создаем счётчик до 1000
// Используем mutex для того чтобы не потерять данные и счётчик ожидания для синхронизации выполнения горутины
func (c *counter) init() int {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Lock()
			c.num++
			c.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return c.num
}
