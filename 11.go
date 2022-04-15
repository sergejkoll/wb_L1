package main

import "fmt"

func main() {
	set1 := make(map[int]struct{})
	set1[0] = struct{}{}
	set1[1] = struct{}{}
	set1[2] = struct{}{}
	set1[3] = struct{}{}

	set2 := make(map[int]struct{})
	set2[2] = struct{}{}
	set2[3] = struct{}{}
	set2[4] = struct{}{}
	set2[5] = struct{}{}

	answer := intersect(set1, set2)
	fmt.Println(answer)
}

// функция нахождения пересечения множеств
// проходим в цикле по наименьшему множеству и проверяем его наличие во втором если есть то добавляем в ответ
func intersect(set1, set2 map[int]struct{}) (answer []int) {
	if len(set1) < len(set2) {
		for k, _ := range set1 {
			if _, ok := set2[k]; ok {
				answer = append(answer, k)
			}
		}
	} else {
		for k, _ := range set2 {
			if _, ok := set1[k]; ok {
				answer = append(answer, k)
			}
		}
	}
	return answer
}
