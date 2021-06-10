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

21.06.09 [parallel]
- 'go' routine
- 'channel' variable
- get multi messages with channel variable


- using go routine & channel with HITURL project
*/

package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestedFailed = errors.New("requested error")

type result struct {
	url    string
	status string
}

func main() {

	//results := make(map[string]string)
	c := make(chan result)

	urls := []string{"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/"}

	//index, value
	for _, url := range urls {
		//fmt.Println(url)
		go hitURL(url, c)
	}

} //end of main

//you can do only send
// func hitURL(url string, c chan<- result) {
// 	fmt.Println("Checking:", url)
// 	//c <- result{} //
// 	fmt.Println(<-c)
// }

func hitURL(url string, c chan<- result) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
		//c <- result{url: url, status: status}
	} //else {

	//}
	c <- result{url: url, status: status}
}
