package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	testStr := "abcd"
	fmt.Println(testStr, "-", checkString(testStr))

	testStr = "abCdefAaf"
	fmt.Println(testStr, "-", checkString(testStr))

	testStr = "aabcd"
	fmt.Println(testStr, "-", checkString(testStr))

	testStr = "abcd"
	fmt.Println(testStr, "-", checkStringWithoutSorting(testStr))

	testStr = "abCdefAaf"
	fmt.Println(testStr, "-", checkStringWithoutSorting(testStr))

	testStr = "aabcd"
	fmt.Println(testStr, "-", checkStringWithoutSorting(testStr))
}

// str -> []runes, сортируем массив и проходим по массиву проверяя соседние
func checkString(str string) bool {
	str = strings.ToLower(str)
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == runes[i+1] {
			return false
		}
	}

	return true
}

// второй вариант с использованием мапы для проверки уникальных символов
func checkStringWithoutSorting(str string) bool {
	str = strings.ToLower(str)
	runes := []rune(str)
	unique := make(map[rune]struct{})

	for _, char := range runes {
		if _, ok := unique[char]; ok {
			return false
		}
		unique[char] = struct{}{}
	}

	return true
}
