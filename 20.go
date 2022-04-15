package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}

// разбиваем строку по пробелу на массив слов
// так же как в 19 меняем их местами и затем объединяем обратно в строку через пробелы
func reverseWords(str string) string {
	s := strings.Fields(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, " ")
}
