package main

import (
	"fmt"

	"github.com/ndgndg91/go-study/banking"
	"github.com/ndgndg91/go-study/facade"
)

func main() {
	bank()
}

func syntaxPractice() {
	facade.HelloGo()
	facade.Variable()
	facade.Functions()
	facade.IfStatement()
	facade.SwitchStatement()
	facade.Pointer()
	facade.ArrayAndSlice()
	facade.MapFunc()
	facade.StructFunc()
}

func bank() {
	account := banking.NewAccount("남동길")
	fmt.Println(account)
}
