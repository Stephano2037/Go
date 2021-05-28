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

*/

package main

import "fmt"

func main() {
	// a := 2
	// b := a
	// a = 10
	// a = 2000000003
	// fmt.Println(&a, &b)

	a := 2
	b := &a
	a = 5
	*b = 20
	fmt.Println(a, *b) //like c pointer
}
