package channel

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var waitGroup sync.WaitGroup
var startTime = time.Now()

func MakeCar() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	waitGroup.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	waitGroup.Wait()
	fmt.Println("Close the Factory")
}

// MakeBody 자동차 차체 생산
func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports Car"
			tireCh <- car
		case <-after:
			close(tireCh)
			waitGroup.Done()
			return
		}
	}
}

// InstallTire 바퀴 생산
func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter Tire"
		paintCh <- car
	}
	waitGroup.Done()
	close(paintCh)
}

// PaintCar 색상 적용
func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car : %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	waitGroup.Done()
}
