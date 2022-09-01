package main

import "fmt"

func main() {

	queue := make(chan string)
	go func() {
		queue <- "one"
		queue <- "two"
		close(queue)
	}()
	for v := range queue {
		fmt.Println(v)
	}
}
