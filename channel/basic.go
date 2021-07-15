package channel

import (
	"fmt"
	"sync"
	"time"
)

func BasicDrive() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch // 9가 들어올 때 까지 기다린다.

	time.Sleep(time.Second)
	fmt.Printf("Square : %d\n", n*n)
	wg.Done()
}
