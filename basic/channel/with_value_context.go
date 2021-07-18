package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wwg sync.WaitGroup

func WithValue() {
	wwg.Add(1)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second) // 5초 후 작업을 종료하는 컨텍스트 생성
	ctx = context.WithValue(ctx, "length", 9)                                   // 컨텍스트릉 래핑하여 key, value 할당

	go printSquare(ctx)
	wwg.Wait()
	cancelFunc()
}

func printSquare(ctx context.Context) {
	if v := ctx.Value("length"); v != nil {
		l := v.(int)
		fmt.Printf("%d \n", l*l)
	}

	wwg.Done()
}
