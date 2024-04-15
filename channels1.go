package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)
	defer close(messages)

	go func() {
		fmt.Println("before ping")
		time.Sleep(time.Second * 3)
		if len(messages) == 0 {
			messages <- "ping"
		} else {
			fmt.Println("channel in func is not empty")
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Wating...")
	fmt.Println(<-messages)

	//channel buffering
	messages2 := make(chan string, 2)
	defer close(messages2)

	messages2 <- "buffered"
	messages2 <- "channel"

	fmt.Println(<-messages2)
	fmt.Println(<-messages2)
}
