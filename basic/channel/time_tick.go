package channel

import (
	"fmt"
	"sync"
	"time"
)

func tickSquare(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            // 1초 간격으로 신호를 보내주는 채널
	terminate := time.After(10 * time.Second) // 10 초 후에 신홀을 보내주는 채널

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Termincated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func TickDrive() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)

	go tickSquare(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	wg.Wait()
}
