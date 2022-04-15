package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, temperature := range temperatures {
		group := int(temperature/10) * 10                  // определяем к какой группе будет принадлежать число
		groups[group] = append(groups[group], temperature) // добавляем число в группу
	}

	for k, v := range groups {
		fmt.Println("Group: ", k, " Value: ", v) // вывод всех групп
	}
}
