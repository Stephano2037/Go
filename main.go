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
*/

package main

import "fmt"

func main() {
	names := []string{"nico", "stephano", "good"}
	names = append(names, "gggggg")
	// names[3] = "get"
	// names[4] = "go"

	fmt.Println(names)
} //end of main
