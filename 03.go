package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// аналогично заданию 2 для каждого числа вызываем горутину которая посчитает квадрат и сложит с atomic переменной
func main() {
	arr := []uint32{2, 4, 6, 8, 10}
	var answer uint32
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, num := range arr {
		go func(number uint32) {
			defer wg.Done()
			atomic.AddUint32(&answer, number*number)
		}(num)
	}

	wg.Wait()
	fmt.Println(answer)
}
