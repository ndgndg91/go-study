package goroutine

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Printf("%d 부터 %d 까지의 합계는 %d 입니다.\n", a, b, sum)
	wg.Done()
}

func Drive() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go sumAtoB(1, 1000000)
	}

	wg.Wait()
}
