package goroutine

import (
	"fmt"
	"time"
)

// GoRoutineWrong : during 5 seconds, at the same time print "ndg", "철수" is sexy %d
// "ndg" Can't print 5 times
func goRoutineWrong() {
	// go main function is not wating for go routine
	// main function end -> go routine die
	go sexyCount("ndg") // this function exit when "ndg is sexy 5"
	sexyCount5("철수")    // because this function end, main function also end.
}

// GoRoutineRight : during 10 seconds, at the same time print "ndg", "철수" is sexy %d
func goRoutineRight() {
	go sexyCount("ndg")
	sexyCount("철수")
}

// during 10 seconds, print 10 times "name is sexy %d"
func sexyCount(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, " is sexy ", i)
		time.Sleep(time.Second)
	}
}

func sexyCount5(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, " is sexy ", i)
		time.Sleep(time.Second)
	}
}
