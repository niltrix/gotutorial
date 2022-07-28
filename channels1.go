package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	//channel buffering
	messages2 := make(chan string, 2)

	messages2 <- "buffered"
	messages2 <- "channel"

	fmt.Println(<-messages2)
	fmt.Println(<-messages2)
}
