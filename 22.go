package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(25_000_000)
	b := big.NewInt(5_000_000)

	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Println("mul: ", mul)

	div := new(big.Int)
	div.Div(a, b)
	fmt.Println("div: ", div)

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Println("sum: ", sum)

	sub := new(big.Int)
	sub.Sub(a, b)
	fmt.Println("sub: ", sub)
}
