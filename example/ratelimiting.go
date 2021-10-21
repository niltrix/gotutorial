package example

import (
	"fmt"
	"time"
)

func RateLimitingtest() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	/*
	 request 는 200 milli 마다 처리됨
	*/
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	/*
		최초 3개의 리미터 할당 후 200 milli 마다 하나씩 할당시도
	*/
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	/*
		5개중 3개는 burstyLimiter가 3개 있어 바로 처리되지만
		2개는 burstyLimiter가 다시 할당되는 주기인 200 milli 마다 처리 된다
	*/
	for req := range burstyRequests {
		<-burstyLimiter /* 할당된것이 있으면 넘어가나 없으면 할당되기까지 대기  */
		fmt.Println("request", req, time.Now())
	}
}
