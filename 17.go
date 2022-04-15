package main

import "fmt"

func main() {
	result := binarySearch([]int{1, 2, 3, 4, 5, 6, 8}, 5)
	fmt.Println(result)
}

func binarySearch(arr []int, key int) bool {
	left := 0
	right := len(arr) - 1
	if arr[left] > key || arr[right] < key {
		return false // если ключа нет в массиве
	}
	for left <= right {
		// берем серидину отсортированного массива
		// если элемент который ищем совпал выход
		// если ключ меньше чем mid поиск в левой части подмассива (left, mid)
		// если ключ больше чем mid поиск в правой части подмассива (mid, right)
		// и так пока не найдем элемент либо пока границы l и r не совпадут (т.е. элемента нет)
		mid := (left + right) / 2
		if arr[mid] == key {
			return true
		}
		if arr[mid] < key {
			left = mid + 1
			continue
		}
		if arr[mid] > key {
			right = mid - 1
			continue
		}
	}
	return false
}
