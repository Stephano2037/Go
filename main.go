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

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	//make channel for get result of 'goroutine' function
	c := make(chan bool)
	people := [2]string{"stephano", "tenten"}
	for _, person := range people {
		//can't return result
		//result := go isSexy(person) -> make channel
		//sending 2 true results through the channel

		go isSexy(person, c)
	}
	//result :=      //wait until it gets just one message
	fmt.Println(<-c) //get one true
	fmt.Println(<-c) //get two true

	// fmt.Println("*******") //district line
	// fmt.Println(<-c)       //get three true ,NOT OPERATING cause it doesn't has message
	// fmt.Println(<-c)       //get four true ,NOT OPERATING

	//main is not receiveing message when you using time sleep
	//time.Sleep(time.Second * 10)
} //end of main

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, channel chan bool) {
	time.Sleep(time.Second * 5)
	//sending true value, through channel variable
	fmt.Println(person)
	channel <- true
	// return true
}
