package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var cwg sync.WaitGroup

func CancelContext() {
	cwg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second) // 작업 시간 3초 동안만 실행
	//ctx, cancelFunc := context.WithCancel(context.Background())
	go printEverySecond(ctx)
	//time.Sleep(5 * time.Second)
	//cancelFunc()

	cwg.Wait()
}

func printEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			cwg.Done()
			return
		case <-tick:
			fmt.Println("Tick!")
		}
	}
}
