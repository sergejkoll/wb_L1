package main

import (
	"fmt"
	"sync"
	"time"
)

var arr = []int{2, 4, 6, 8, 10}

// первое самое простое решение
// для каждого числа запускаем горутину в которой выводим квадрат числа
// sleep для ожидания всех горутин
func firstSolution() {
	for _, num := range arr {
		go func(number int) {
			fmt.Println(number * number)
		}(num)
	}

	time.Sleep(time.Second * 2)
}

// второе решение
// используем небуферезированный канал
// для каждого числа из массива запускаем горутину которая пишет в канал значение квадрта числа
// после чего читаем из канала len(arr) раз
func secondSolution() {
	answer := make(chan int)
	for i, _ := range arr {
		go func(i int) {
			answer <- arr[i] * arr[i]
		}(i)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(<-answer)
	}
}

// третье решение
// используем waitGroup и буферезированный канал
// также для каждого числа запускаем горутину в которой пишем в буфферизриованный канал квадрат числа
// ждем выполнения всех горутин с помощью wg.Wait(), закрываем канал и в цикле читаем все его значения
func thirdSolution() {
	answer := make(chan int, len(arr))
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, num := range arr {
		go func(number int) {
			defer wg.Done()
			answer <- number * number
		}(num)
	}

	wg.Wait()
	close(answer)

	for num := range answer {
		fmt.Println(num)
	}

}

func main() {
	firstSolution()
	secondSolution()
	thirdSolution()
}
