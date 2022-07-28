package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func Coroutinetest() {

	//Parallel processing for Logical CPUs
	runtime.GOMAXPROCS(4)

	f("direct")

	go f("goroutine")

	go func(msg string) {
		time.Sleep(time.Second * 1)
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 2)
	fmt.Println("done")

	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		time.Sleep(time.Second * 5)
		defer wait.Done()
		fmt.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done()
		time.Sleep(time.Second * 2)
		fmt.Println(msg)
	}("Hi")

	fmt.Println("Wait coroutine done")
	wait.Wait()
	fmt.Println("Coroutine done")
}
