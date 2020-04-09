package main

import (
	"fmt"
	"sync"
	"time"
)

func helloWithNumber(in int) {
	fmt.Println("Hej: ", in)
}

func takeStringPutOnChannel(someString string, message chan string) {
	message <- someString
}

func hej(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	helloWithNumber(in)
}

func sendGrej(ch chan int) {
	ch <- 1
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go hej(1, &wg)
	go hej(2, &wg)
	go hej(3, &wg)
	wg.Wait()

	// Test of channel
	message := make(chan string)

	go takeStringPutOnChannel("first ", message)
	go takeStringPutOnChannel("second ", message)

	info1, info2 := <-message, <-message

	fmt.Println(info1, " and ", info2)

}
