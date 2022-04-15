package main

import "fmt"

type Human struct {
	name             string
	age              int
	accumulatedMoney float64
}

func (h Human) GetName() string {
	return h.name
}

func (h Human) GetAge() int {
	return h.age
}

func (h Human) GetAccumulatedMoney() float64 {
	return h.accumulatedMoney
}

type Action struct {
	Human
	actionType string
}

func (a *Action) AddMoney(amount float64) {
	a.accumulatedMoney += amount
}

func main() {
	human := Human{name: "Alex", age: 21}
	action := Action{Human: human, actionType: "work"}
	action.AddMoney(100)

	fmt.Println(human.GetName())
	fmt.Println(human.GetAge())
	fmt.Println(action.GetAccumulatedMoney())
}
