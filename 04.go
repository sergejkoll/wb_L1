package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	ch := make(chan interface{})                 // главный канал
	signalChannel := make(chan os.Signal, 1)     // канал для обработки SIGINT
	signal.Notify(signalChannel, syscall.SIGINT) // запись в signalChannel если пришел SIGINT

	// запуск N воркеров
	for i := 0; i < N; i++ {
		go worker(ch)
	}

	// бесокнечная запись в канал и обрабтка signalChannel
	for i := 0; ; i++ {
		select {
		case <-signalChannel:
			fmt.Println("Завершение работы")
			close(ch)
			close(signalChannel)
			return
		default:
			ch <- i
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func worker(ch chan interface{}) {
	// 1 варинат \\
	//for {
	//	item, ok := <-ch
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(item)
	//}
	for item := range ch {
		fmt.Println(item)
	}
	fmt.Println("Завершение работы горутины")
}
