package channel

import (
	"fmt"
	"sync"
	"time"
)

func waitingSquare(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // 채널에서 데이터를 계속 기다림.
		fmt.Printf("Sqaure : %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done() // 실행되지 않는다.
}

func WaitingAndHandleAndDeadLock() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go waitingSquare(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2 // 채널에 데이터를 넣는다. 0, 2, 4, 6, 8 ... 18
	}
	wg.Wait() // 작업 완료를 기다리지만 waitingSquare 함수에서 wg.Done 이 실행되지 않아 작업이 완료되지 않는다.
}
