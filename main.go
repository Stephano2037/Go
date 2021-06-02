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

*/

package main

import (
	"chanhyun/banking"
	"fmt"
)

func main() {
	// account := banking.Account{Owner: "stephano", Balance: 1000}
	// account.Owner = "Coke"

	//account := banking.Account
	account := banking.NewAccount("Stephano")
	account.Deposit(100)
	fmt.Println(account.Balance())
	//fmt.Println(account) //return object (not copy)
} //end of main
