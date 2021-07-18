package syntax

import (
	"fmt"
	"strings"
)

// Multiply : a * b
func Multiply(a int, b int) int {
	return a * b
}

// RepeatMe : words array print
func RepeatMe(words ...string) {
	fmt.Println(words)
}

// LenAndUpper : length, uppercase
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// LenAndUpperReturnVariable : length, uppercase using naked return, defer
func LenAndUpperReturnVariable(name string) (length int, upper string) { // naked return
	defer fmt.Println("lenAndUpperReturnVariable is Done.") // after the function, defer clause trigger.
	length = len(name)
	upper = strings.ToUpper(name)
	return
}
