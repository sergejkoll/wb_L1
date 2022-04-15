package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{}) // итоговое множество слов

	for _, word := range words {
		if _, ok := set[word]; !ok { // если слова еще нет в множестве то добавляем его
			set[word] = struct{}{}
		}
	}

	for k, _ := range set {
		fmt.Println(k)
	}
}
