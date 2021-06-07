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

*/

package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestedFailed = errors.New("requested error")

func main() {
	urls := []string{"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/"}

	//index, value
	for _, url := range urls {
		//fmt.Println(url)
		hitURL(url)
	}

} //end of main

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestedFailed
	}
	//100,200,300 -> redirection, 400 series-> error
	return nil
}
