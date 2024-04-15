package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe")
	ann := boring("Ann")

	for i := 0; i < 5; i++ {
		// 각 블락킹 모드이기에 먼저 도착한 메시지가 먼저 출력되는것이 아닌 항상 아래 순서대로 출력된다
		// ann 이 메세지가 먼저 도착하더라도 joe 메세지가 도착한 이후 실행되어 동시성에 좋지 않다.
		fmt.Printf("You say: %q\n", <-joe) // Blocking point!! Receive expression is just a value.
		fmt.Printf("You say: %q\n", <-ann) // Blocking point!! Receive expression is just a value.
	}
	fmt.Println("You're boring; I'm leaving.")
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
