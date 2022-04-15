package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("before sleep")
	sleep(time.Second * 2)
	fmt.Println("after sleep")
}

func sleep(n time.Duration) {
	<-time.After(n)
}
