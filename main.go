package main

import (
	"fmt"
	"log"

	"github.com/ndgndg91/go-study/banking"
	"github.com/ndgndg91/go-study/facade"
	"github.com/ndgndg91/go-study/mydict"
)

func main() {
	dictionary()
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
	account.Deposit(100)
	fmt.Println(account.Balance())
	fmt.Println(account)

	err := account.Withdraw(200) // this code have to handle error
	if err != nil {
		log.Fatalln(err) // this kill the program
	}

	fmt.Println(account)
}

func dictionary() {
	dictionary := mydict.Dictionary{"first": "First Word"}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)

	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	addErr := dictionary.Add("hello", "Greeting")
	if addErr != nil {
		fmt.Println("addErr", addErr)
	}

	definition, helloSearchErr := dictionary.Search("hello")
	if helloSearchErr != nil {
		fmt.Println(helloSearchErr)
	} else {
		fmt.Println(definition)
	}
}
