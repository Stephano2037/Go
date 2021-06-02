package banking

import (
	"errors"
	"fmt"
)

//private struct
type bankAccount struct {
	owner   string
	balance int
}

//public struct
type Account struct {
	//Upper case for public variable in struct
	// Owner   string
	// Balance int

	owner   string
	balance int
}

//google recommend to insert "err" name,
// in front of custom error variable
var errNomoney = errors.New("Can't withdraw")

// newaccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Now struct has this function 가지고 있음 (receiver function)
// like a class
// Must type lower case of Struct First alphabet (ex: 'A'ccount -> a)
//Deposit x amount on your account
//It makes "copy" of Account object (not itself)
/*
should add * (pointer), real receiver
*/
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//print balance
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {

	if a.balance < amount {
		//go didn't tell us anything about error
		//write should handle these things
		return errNomoney //errors.New("Can't withdraw. You are poor")
	}
	a.balance -= amount

	return nil //same like none or null
}

//change owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//owner of the account
func (a Account) Owner() string {
	return a.owner
}

//Automatically return string in Go (like a c++ class)
func (a Account) String() string {
	//Sprint -> string print
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
