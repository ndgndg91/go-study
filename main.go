package main

import (
	"fmt"

	"github.com/ndgndg91/go-study/greetings"
	"github.com/ndgndg91/go-study/syntax"
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

	fmt.Println(syntax.Multiply(2, 2))

	lenOfName, upperName := syntax.LenAndUpper("ndgndg91")
	fmt.Println(lenOfName, upperName)

	syntax.RepeatMe("a", "b", "c", "d", "e", "f")

	length, upperCase := syntax.LenAndUpperReturnVariable("giri")
	fmt.Println(length, upperCase)

	fmt.Println(syntax.Sum(1, 2, 3, 4, 5))

	fmt.Println(syntax.Sum2(5, 6, 7, 8, 9))
}
