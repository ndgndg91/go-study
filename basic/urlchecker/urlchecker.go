package urlchecker

import (
	"errors"
	"fmt"
	"net/http"
)

func hitURLDriver() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.instagram.com/",
		"https://dong-gil.com/",
	}

	for _, url := range urls {
		result := "OK"
		hitErr := hitURL(url)
		if hitErr != nil {
			result = "FAIL"
		}

		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFail = errors.New("Request failed")

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFail
	}

	return nil
}

type requestResult struct {
	url    string
	status string
}

// HitURLUsingRoutine *
func HitURLUsingRoutine() {
	var results = make(map[int]requestResult)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.instagram.com/",
		"https://dong-gil.com/",
	}

	for _, url := range urls {
		go hitURLUsingChannel(url, c)
	}

	for i := 0; i < len(urls); i++ {
		results[i] = <-c
	}

	for key, value := range results {
		fmt.Println(key, value)
	}
}

// this channel only send message ( chan<- )
func hitURLUsingChannel(url string, c chan<- requestResult) {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- requestResult{url: url, status: status}
}
