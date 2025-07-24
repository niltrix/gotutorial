package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(intput1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-intput1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg string) <-chan string { // Returns recive-only channel of strings.
	c := make(chan string)
	go func() { // Launch go routine to generate the message.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %q\n", <-c) // Blocking point!! Receive expression is just a value.
	}
	fmt.Println("You're boring; I'm leaving.")
}
