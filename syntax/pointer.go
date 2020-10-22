package syntax

import "fmt"

// PrintAddress : print memory address, value
func PrintAddress() {
	var a int = 10
	b := 5

	fmt.Printf("a value : %d, b value : %d \n a address : %d, b address : %d\n", a, b, &a, &b)
}

// PrintPointMemory : point memory address
func PrintPointMemory() {
	var a int = 20
	b := &a

	// *b see through a value => a value
	// &b memory address of variable b
	fmt.Printf("a value : %d,  a address : %d, b value : %d, b address : %d, b see through : %d\n", a, &a, b, &b, *b)

	a = 30
	fmt.Printf("a value : %d,  a address : %d, b value : %d, b address : %d, b see through : %d\n", a, &a, b, &b, *b)

	*b = 100
	fmt.Printf("a value : %d,  a address : %d, b value : %d, b address : %d, b see through : %d\n", a, &a, b, &b, *b)
}
