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
*/

package main

import "fmt"

type person struct {
	name    string
	age     int
	favfood []string
}

func main() {
	favFood := []string{"kimchi", "pizza"}
	//stephano := person{"Stephano", 29, favFood}
	stephano := person{name: "stephano", age: 18, favfood: favFood}

	fmt.Println(stephano)
} //end of main
