package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	msgChannel := make(chan int, 3)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go receiver(msgChannel, &wg)
	for i := 1; i <= 6; i++ {
		size := len(msgChannel)
		fmt.Printf("%s Sending: %d. Buffer Size: %d\n",
			time.Now().Format("15:04:05"), i, size)
		msgChannel <- i
	}
	msgChannel <- -1
	wg.Wait()
}

func receiver(messages chan int, wg *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = <-messages
		fmt.Println(time.Now().Format("15:04:05"), "Received:", msg)
	}
	wg.Done()
}
