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

*/

package main

import (
	"chanhyun/mydict"
	"fmt"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	fmt.Println(dictionary)
	err := dictionary.Update("hello", "Second")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	err = dictionary.Delete("g")

	fmt.Println(err, dictionary)

	//dictionary := mydict.Dictionary{"first": "First word"}
	//definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// err = dictionary.Add("second", "Greeting")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// definition, err = dictionary.Search("second")

	// fmt.Println(definition, err)
	// dictionary := mydict.Dictionary{}

	// dictionary["hello"] = "Bye"
	// fmt.Println(dictionary)
} //end of main
