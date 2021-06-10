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
	//go routine is available when main is running itself
	go sexyCount("stephano")

	//if you put the go word this function, main function will be closed first
	//main is counting sexycount of tenten
	sexyCount("tenten")
	//time.Sleep(time.Second * 5)// alive just 5 seconds
	//go sexycount -> main is finished, it doesn't wait for another function

} //end of main

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
