package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// если не передали аргумент
	if len(os.Args) < 2 {
		fmt.Println("missing command line argument")
		return
	}

	// проверка полученного аргумента
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid command line argument")
		return
	}

	ch := make(chan interface{})                        // канал в который будем последовательно писать и читать
	timer := time.After(time.Duration(N) * time.Second) // таймер который используется для остановки producer
	var wg = sync.WaitGroup{}                           // для ожидания завершения горутины consumer
	wg.Add(1)

	// producer
	go func() {
		for i := 0; ; i++ {
			select {
			case <-timer:
				fmt.Println("Завершение работы горутины producer")
				close(ch)
				return
			default:
				ch <- i
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	// consumer
	go func() {
		defer wg.Done()
		for item := range ch {
			fmt.Println(item)
		}
		fmt.Println("Завершение работы горутины consumer")
	}()

	wg.Wait()
}
