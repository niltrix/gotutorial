package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			c1 <- "one"
		}
	}()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			c2 <- "two"
		}
	}()

	fmt.Println("start select------------------")
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("     2s received", msg1)
		case msg2 := <-c2:
			fmt.Println("        5s received", msg2)
		default:
			{
				time.Sleep(1 * time.Second)
				fmt.Println("default pass")
			}
		}
	}
	fmt.Println("end select-------------------\n\n")
}
