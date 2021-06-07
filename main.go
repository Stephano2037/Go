/*
21.05.23
Learning GO language with nomad code

Using simple print function  (Hello World!)


add something package (21.05.23)


learning const variable & nomral variable (21.05.23)

add comment for github third (21.05.23)


21.05.27 Functions part One
--Naked Function
-- Defer

21.05.27 for loop
- range


21.05.29
- if else | switch
- Pointer in GO

21.05.31
- Array and slices
- Map
- Struct


21.06.02 [Bank & Dictionary projects]
- Account + NewAccount


21.06.03 [Dictionary]


21.06.07  [URL checker project]
- url check (url hit)
- using map
*/

package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestedFailed = errors.New("requested error")

func main() {

	//Cause Error , but compiler doesn't know about this
	//var results map[string]string

	//var results := map[string]string{}

	//make is making map with empty
	var results = make(map[string]string) //nil

	urls := []string{"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/"}

	results["bye"] = "hello"
	//index, value
	for _, url := range urls {
		result := "ok"
		//fmt.Println(url)
		err := hitURL(url)
		if err != nil {
			result = "failed"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}

	//fmt.Println(results)

} //end of main

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		//if i want to see about each site error
		fmt.Println(err, resp.StatusCode)
		return errRequestedFailed
	}
	//401 : unauthorized
	//429 : too many requests
	//100,200,300 -> redirection, 400 series-> error
	return nil
}
