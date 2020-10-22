package main

import (
	"fmt"
	"strings"

	"github.com/ndgndg91/go-study/greetings"
)

func main() {
	fmt.Println("hello go!")
	greetings.SayHi()
	greetings.SayBye()

	const name string = "남동길"
	// name = "김철수" 상수는 변경이 불가능
	fmt.Println(name)

	var fullName string = "남동길"
	address := "서울시 중구 퇴계로 235 남산센트럴자이" // type inference -> address is string, this syntax only apply in function scope
	fullName = "김철수"

	address = "서울시 마포구 홍대입구"

	fmt.Println(fullName)
	fmt.Println(address)

	fmt.Println(multiply(2, 2))

	lenOfName, upperName := lenAndUpper("ndgndg91")
	fmt.Println(lenOfName, upperName)

	repeatMe("a", "b", "c", "d", "e", "f")

	length, upperCase := lenAndUpperReturnVariable("giri")
	fmt.Println(length, upperCase)

	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println(sum2(5, 6, 7, 8, 9))
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpperReturnVariable(name string) (length int, upper string) { // naked return
	defer fmt.Println("lenAndUpperReturnVariable is Done.") // after the function, defer clause trigger.
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func multiply(a int, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers { // '_' -> ignore
		total += number
	}

	return total
}

func sum2(numbers ...int) int {
	var total int = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	return total
}
