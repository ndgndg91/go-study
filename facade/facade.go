package facade

import (
	"fmt"

	"github.com/ndgndg91/go-study/greetings"
	"github.com/ndgndg91/go-study/syntax"
)

// HelloGo : hello Go!
func HelloGo() {
	fmt.Println("hello go!")
	greetings.SayHi()
	greetings.SayBye()
}

// Variable : practice variable
func Variable() {
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

// Functions : practice function
func Functions() {
	fmt.Println(syntax.Multiply(2, 2))

	lenOfName, upperName := syntax.LenAndUpper("ndgndg91")
	fmt.Println(lenOfName, upperName)

	syntax.RepeatMe("a", "b", "c", "d", "e", "f")

	length, upperCase := syntax.LenAndUpperReturnVariable("giri")
	fmt.Println(length, upperCase)

	fmt.Println(syntax.Sum(1, 2, 3, 4, 5))

	fmt.Println(syntax.Sum2(5, 6, 7, 8, 9))
}

// IfStatement : practice if
func IfStatement() {
	fmt.Println(syntax.CanIDrink(20))
}

// SwitchStatement : practice switch
func SwitchStatement() {
	fmt.Println(syntax.WhatIsGrade(91))
}

// Pointer : practice pointer
func Pointer() {
	syntax.PrintAddress()
	syntax.PrintPointMemory()
}

// ArrayAndSlice : practice array and slice
func ArrayAndSlice() {
	names := syntax.GetNames()
	fmt.Println(names)
	// names = append(names, "슬라이스 불가능")
	var sliceArray []string = syntax.SliceTest()
	fmt.Println(sliceArray)
	sliceArray = append(sliceArray, "sliceTest")
	fmt.Println(sliceArray)
}

// MapFunc : practice map
func MapFunc() {
	var myMap = syntax.NewMap()
	myMap[100] = "paul"
	fmt.Println(myMap)
	myMap[1] = "덮어쓰기"
	fmt.Println(myMap)

	syntax.AddToMap(myMap, 50, "내가 만든 function")
	syntax.PrintMapUsingRange(myMap)
}

// StructFunc : practice struct
func StructFunc() {
	interests := []string{"k8s", "go", "java", "spring", "aws"}
	ndg := syntax.CreateUsers("남동길", 30, interests)
	fmt.Println(ndg)
}
