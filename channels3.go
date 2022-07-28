package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("Received message ", msg)
	default:
		fmt.Println("no message received")

	}

	msg := "hi"

	select {
	case message <- msg:
		fmt.Println("Send message ", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message:
		fmt.Println("Receivced message ", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	// Timeout pattern example
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
