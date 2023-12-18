package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan interface{})

	ctx, cancel := context.WithCancel(context.Background())

	// Канал для записи сигнала
	exitChannel := make(chan os.Signal, 1)

	// Генерация Ctrl+C в канал
	signal.Notify(exitChannel, syscall.SIGINT)

	g, ctx := errgroup.WithContext(ctx)
	const N = 5
	for i := 0; i < N; i++ {
		workerNum := i
		g.Go(func() error {
			return worker(workerNum, ch, ctx)
		})
	}

	writerChannel(exitChannel, ch)
	cancel()
	err := g.Wait()
	if err != nil {
		fmt.Printf("got error from workers: %s", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func writerChannel(exitCH chan os.Signal, ch chan interface{}) {
	for {
		select {
		case <-exitCH:
			fmt.Println("gracefully shutdown from writer")
			return
		default:
			ch <- 42
			time.Sleep(time.Nanosecond)
		}
	}
}

func worker(id int, ch chan interface{}, ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("gracefully shutdown exit from worker")
			return nil
		case data := <-ch:
			fmt.Printf("Работник %d получил данные: %v\n", id, data)
			time.Sleep(1 * time.Second)
		default:
		}
	}
}
