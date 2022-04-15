package main

import "fmt"

func main() {
	var number int64
	number = 15
	fmt.Println(number) // 1111

	i := 4                           // (1 << 4) == 10000
	numberWith1 := number | (1 << i) // 01111 | 10000 = 11111
	fmt.Println(numberWith1)         // 11111 = 31

	i = 2                             // (1 << 2) == 100
	numberWith0 := number &^ (1 << i) // 1111 & !(0100) == 1111 & 1011 = 1011
	fmt.Println(numberWith0)          // 1011 = 11
}
