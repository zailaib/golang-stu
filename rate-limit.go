package main

import "fmt"
import "time"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limiter
		fmt.Println("request ", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	bursyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		bursyRequests <- i
	}
	close(bursyRequests)

	for req := range bursyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
