package main

import (
	"fmt"

	"github.com/ndgndg91/go-study/greetings"
	"github.com/ndgndg91/go-study/syntax"
)

func main() {

}

func helloGo() {
	fmt.Println("hello go!")
	greetings.SayHi()
	greetings.SayBye()
}

func variable() {
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

func functions() {
	fmt.Println(syntax.Multiply(2, 2))

	lenOfName, upperName := syntax.LenAndUpper("ndgndg91")
	fmt.Println(lenOfName, upperName)

	syntax.RepeatMe("a", "b", "c", "d", "e", "f")

	length, upperCase := syntax.LenAndUpperReturnVariable("giri")
	fmt.Println(length, upperCase)

	fmt.Println(syntax.Sum(1, 2, 3, 4, 5))

	fmt.Println(syntax.Sum2(5, 6, 7, 8, 9))
}

func ifStatement() {
	fmt.Println(syntax.CanIDrink(20))
}

func switchStatement() {
	fmt.Println(syntax.WhatIsGrade(91))
}

func pointer() {
	syntax.PrintAddress()
	syntax.PrintPointMemory()
}

func arrayAndSlice() {
	names := syntax.GetNames()
	fmt.Println(names)
	// names = append(names, "슬라이스 불가능")
	var sliceArray []string = syntax.SliceTest()
	fmt.Println(sliceArray)
	sliceArray = append(sliceArray, "sliceTest")
	fmt.Println(sliceArray)
}

func mapFunc() {
	var myMap = syntax.NewMap()
	myMap[100] = "paul"
	fmt.Println(myMap)
	myMap[1] = "덮어쓰기"
	fmt.Println(myMap)

	syntax.AddToMap(myMap, 50, "내가 만든 function")
	syntax.PrintMapUsingRange(myMap)
}
