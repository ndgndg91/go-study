package main

import (
	"fmt"

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
}
