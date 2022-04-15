package main

import "fmt"

// UserService - Целевой интерфейс
type UserService interface {
	Request() string
}

// Тип который необходимо адаптировать
type admin struct {
}

// NewAdapter - конструктор адаптера
func NewAdapter(a *admin) UserService {
	return &Adapter{a}
}

// SpecificRequest - не адаптированный метод
func (a *admin) SpecificRequest() string {
	return "Request"
}

// Adapter - адаптер реализующий целевой интерфейс
type Adapter struct {
	*admin
}

// Request - адаптированный метод
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	a := &admin{}
	adapter := NewAdapter(a)
	fmt.Println(adapter.Request())
}
