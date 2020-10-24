package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"

	"github.com/ndgndg91/go-study/banking"
	"github.com/ndgndg91/go-study/channel"
	"github.com/ndgndg91/go-study/facade"
	"github.com/ndgndg91/go-study/mydict"
	"github.com/ndgndg91/go-study/urlchecker"
)

var baseURL string = "https://kr.indeed.com/jobs?q=java&limit=50"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(pageNumber int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(pageNumber*50)
	fmt.Println("Requesting", pageURL)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", res.StatusCode)
	}
}

func urlChecker() {
	urlchecker.HitURLUsingRoutine()
}

func goRoutineAndChannel() {
	channel.ChannelPractice()
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

	updateErr := dictionary.Update("hello", "Updated Greeting")
	if updateErr != nil {
		fmt.Println(updateErr)
	}

	updatedDefinition, afterUpdatedSearchErr := dictionary.Search("hello")
	if afterUpdatedSearchErr != nil {
		fmt.Println(afterUpdatedSearchErr)
	} else {
		fmt.Println(updatedDefinition)
	}

	cantUpdateErr := dictionary.Update("non-exists", "try update")
	if cantUpdateErr != nil {
		fmt.Println(cantUpdateErr)
	}

	dictionary.Delete("hello")
	fmt.Println(dictionary)
}
