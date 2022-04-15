package main

import (
	"fmt"
	"strconv"
	"sync"
)

type storage struct {
	sync.RWMutex
	data map[int]string
}

// или использовать sync.Map

func (s *storage) Set(key int, value string) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}

func (s *storage) Get(key int) string {
	s.RLock()
	defer s.RLock()
	return s.data[key]
}

func main() {
	data := storage{data: map[int]string{}}
	data.Set(1, "1")

	var wg sync.WaitGroup // для примера, чтобы дождаться выполнения всех горутин
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(item int) {
			defer wg.Done()
			data.Set(item, strconv.Itoa(item))
		}(i)
	}
	wg.Wait()
	fmt.Println(data.Get(0))
	fmt.Println(data.Get(1))
	fmt.Println(data.Get(2))
	fmt.Println(data.Get(3))
	fmt.Println(data.Get(4))
}
