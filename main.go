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
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	//make channel for get result of 'goroutine' function
	c := make(chan string)
	people := [5]string{"stephano", "tenten", "dal", "hoho", "abc"}
	for _, person := range people {
		//can't return result
		//result := go isSexy(person) -> make channel
		//sending 2 true results through the channel

		go isSexy(person, c)

	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

	/*
		fmt.Println("waiting for message")
		resultOne := <-c
		resultTwo := <-c
		//Remember it is not "procedurellY operating"
		fmt.Println("Received this message: ", resultOne) //get one true , waiting for a message until we get another message
		fmt.Println("Received this message: ", resultTwo) //get two true
	*/
	//time.Sleep(time.Second * 2)
} //end of main

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	//sending true value, through channel variable
	//fmt.Println(person)

	channel <- person + " is sexy"
	// return true
}
