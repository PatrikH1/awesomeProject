package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 5 {
		time.Sleep(4 * time.Second)
		fmt.Println("hello world")
		i++
	}
}
