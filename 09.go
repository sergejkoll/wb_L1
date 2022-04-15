package main

import "fmt"

// конвейер чисел
func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := gen(arr) // ch - канал для записи чисел из массива
	out := sq(ch)  // out - канал для записи квадратов чисел

	for n := range out {
		fmt.Println(n) // вывод в stdout квадратов чисел
	}
}

func gen(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
