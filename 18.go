package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	sync.Mutex
	Number int
}

func (c *counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.Number += 1
}

type atomicCounter struct {
	Number uint64
}

func (ac *atomicCounter) Increment() {
	atomic.AddUint64(&ac.Number, 1)
}

func main() {
	// используем структуру с мьютексом в которой хранится наш счетчик
	// при каждой операции с числом блокируем мьютекс
	c := counter{
		Mutex:  sync.Mutex{},
		Number: 0,
	}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				c.Increment()
			}
		}()
	}
	wg.Wait()

	fmt.Println(c.Number)

	// 2.
	// использование атомарной переменной вместо мьютексов
	ac := atomicCounter{Number: 0}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				ac.Increment()
			}
		}()
	}
	wg.Wait()

	fmt.Println(ac.Number)
}
