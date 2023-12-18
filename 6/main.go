package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1) Естественное завершение программы c использованием каналов.
	//var wg sync.WaitGroup
	//ch := make(chan string)
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	ch <- "hello"
	//	close(ch)
	//}()
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println(<-ch)
	//}()
	//wg.Wait()
	//time.Sleep(time.Second)

	//2) При помощи логического флага через каналы.
	//ch := make(chan bool)
	//flag := make(chan bool)
	//go func() {
	//	for {
	//		select {
	//		case <-flag:
	//			fmt.Println("Горутина корректно завершает работу")
	//			ch <- true
	//			return
	//		default:
	//			fmt.Println("Выполнение горутины...")
	//			time.Sleep(time.Second)
	//		}
	//	}
	//}()
	//time.Sleep(time.Second * 5)
	//flag <- true
	//fmt.Println("Done", <-ch)

	//3)runtime.Goexit
	//go func() {
	//	fmt.Println("Выполнение горутины..")
	//	runtime.Goexit()
	//}()
	//time.Sleep(time.Second)

	//4)При помощи использование context.Context.
	//Пакет context в Go предоставляет способ распространения сигналов отмены по горутинам.
	//ctx, cancel := context.WithCancel(context.Background())
	////ch := make(chan string)
	//go contextFunc(ctx)
	//time.Sleep(time.Second * 5)
	//cancel()
	//time.Sleep(time.Second)
	//fmt.Println("Done main")
}

func contextFunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context exit...")
			return
		default:
			fmt.Println("Выполнение горутины...")
			time.Sleep(time.Second)
		}
	}
}

//Небуферизованный канал в Go может одновременно хранить
//только один фрагмент данных. После того как значение записано в канал,
//его необходимо прочитать, прежде чем можно будет записать другое значение

//Вывод значения из канала не работает из-за того,
//что операция чтения из канала происходит во второй горутине до того,
//как значение будет записано в канал в первой горутине. Из-за этого получается
//блокировка операции чтения и программа зависает. (Нужна сихронизация)
