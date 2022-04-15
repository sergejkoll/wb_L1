package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8}
	newArr := remove(arr, 2)
	fmt.Println(newArr)
}

// функция удаления элемента
// проверяем переданный индекс
// если все ок - возвращаем новый массив без элемента по этому индексу
func remove(arr []int, idx int) []int {
	if idx < 0 || len(arr) < 0 || idx >= len(arr) {
		return []int{}
	}
	return append(arr[:idx], arr[idx+1:]...) // вторым аргументов с помощью ... передаем список
}
