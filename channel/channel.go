package channel

import (
	"fmt"
	"math/rand"
	"time"
)

// ChannelPractice : main function is waiting for the channel to receive result
func ChannelPractice() {
	c := make(chan string)
	people := [2]string{"ndg", "철수"}

	for _, person := range people {
		go isSexy(person, c) // 2 go routine
	}

	fmt.Println("waiting for receving messages")
	// <-c is blocking operation
	resultFirst := <-c  // waiting for receiving first routine
	resultSecond := <-c // waiting for receiving second routine
	fmt.Println(resultFirst)
	fmt.Println(resultSecond)

	// complier do not know that how many go routine
	// but, on the run time go runner know that below code is not correct
	// so, program die for deadlock

	// resultThird := <-c
	// fmt.Println(resultThird)

	// so, this way is not good
	// use for loop
	// for receiving message
	// like a below code

	// for i := 0; i < len(people); i++ {
	// fmt.Println(<-c)
	// }
}

func isSexy(person string, c chan string) {
	n := rand.Intn(5)
	fmt.Println(n, "second wait..")
	d := time.Duration(n) * time.Second
	time.Sleep(d)
	c <- person + " is sexy" // sending message to channel
}
