package main

import "fmt"

func main() {
	integer := 10
	str := "string"
	boolean := true
	channel := make(chan interface{})

	getType(integer)
	getType(str)
	getType(boolean)
	getType(channel)
}

func getType(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	case bool:
		fmt.Println("bool: ", v)
	case chan interface{}:
		fmt.Println("chan interface{}: ", v)
	}
}
