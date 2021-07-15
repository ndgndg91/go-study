package channel

import "fmt"

// ZombieRoutine 채널에 데이터가 있지만 보관할 곳이 없기 때문에 다른 고루틴에서 데이터를 빼가기를 기다린다.
func ZombieRoutine() {
	ch := make(chan int) // 크기가 0인 채널 생성

	ch <- 9                    //  함수는 여기서 먼춘다.
	fmt.Println("Never print") // 실행되지 않는다.
}
