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
*/

package main

import "fmt"

func main() {
	// "string,string" [key]value
	map_variable := map[string]string{"name": "stephano", "age": "12"} //can't use int
	//fmt.Println(map_variable)
	// for _, value := range map_variable {
	// 	fmt.Println(value)
	// }

	for key, value := range map_variable {
		fmt.Println(key, value)
	}

	for key, _ := range map_variable {
		fmt.Println(key)
	}

} //end of main
