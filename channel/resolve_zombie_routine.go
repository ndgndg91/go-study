package channel

import (
	"fmt"
	"sync"
	"time"
)

func doNotWaitingSquare(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // close(ch) 를 통해서 닫힌 채널을 더이상 대기하지 않는다.
		fmt.Printf("Square : %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func ResolveZombieRoutine() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go doNotWaitingSquare(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 채널을 닫아줌
	wg.Wait()
}
