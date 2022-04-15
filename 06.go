package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1. используя канал для остановки
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Завершение работы горутины")
				close(done)
				return
			default:
				fmt.Println(rand.Int())
			}
		}
	}()

	time.Sleep(time.Second)
	done <- struct{}{}
	time.Sleep(time.Second)

	// 2. используя таймер
	timer := time.NewTimer(time.Second)

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Завершение работы горутины")
				return
			default:
				fmt.Println(rand.Int())
			}
		}
	}()

	time.Sleep(time.Second * 2)

	// 3. Используя контекст
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Завершение работы горутины")
				return
			default:
				fmt.Println(rand.Int())
			}
		}
	}(ctx)

	time.Sleep(time.Second)
	cancel()
}
