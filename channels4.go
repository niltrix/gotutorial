package main

import "fmt"

func main() {
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("작업 수행 : ", j)
			} else {
				fmt.Println("더 이상 작업이 없음")
				done <- true
				return
			}
		}
	}()

	for j := 0; j <= 3; j++ {
		jobs <- j
		fmt.Println("작업 요청 : ", j)
	}
	close(jobs)

	fmt.Println("모든 작업 마침")
	//blocking until receive data from channel
	//wait until go func()
	<-done

	//close 는 channel의 buffer 가 비워져야 비로소 close가 수행됨됨
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
