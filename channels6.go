package main

import (
	"fmt"
	"time"
)

func consuming(scheduler chan string) {
	select {
	case <-scheduler:
		fmt.Println("received..")
	case <-time.After(5 * time.Second):
		fmt.Println("time over")
	}
}

func producing(scheduler chan string) {
	var name string
	fmt.Print("input:")
	fmt.Scanln(&name)
	scheduler <- name
}

func main() {
	scheduler := make(chan string)
	go consuming(scheduler)
	go producing(scheduler)

	time.Sleep(100 * time.Second)
}
