package main

import (
	"fmt"
	"time"
)

func worker1(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 1)
	fmt.Println("done")

	done <- true
}

type job struct {
	Name   string
	Result string
}

func worker2(j chan job) {
	for {
		myjob := <-j
		if myjob.Name == "close" {
			time.Sleep(time.Second * 1)
			myjob.Result = "close"
		} else {
			fmt.Println("working...: ", myjob.Name)
			myjob.Result = "done"
		}
		j <- myjob
	}
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	//basic
	fmt.Println("Basic test : Channel sync  ...")
	done := make(chan bool, 1)
	go worker1(done)
	fmt.Println("main : waiting ...")
	<-done //blocking until receive data from channel

	//advance
	fmt.Println("Advance test : Channel sync  ...")
	jobChannel := make(chan job, 1)
	go worker2(jobChannel)
	for _, i := range []string{"job 1", "job 2"} {
		jobChannel <- job{Name: i}
		result := <-jobChannel
		fmt.Println("Job: ", result.Name, result.Result)
	}
	jobChannel <- job{Name: "close"}
	result := <-jobChannel
	fmt.Println("result: ", result)

	//Channel direction
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
