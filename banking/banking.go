package banking

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

// newaccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Now struct has this function 가지고 있음 (receiver function)
// like a class
// Must type lower case of Struct First alphabet (ex: 'A'ccount -> a)
//Deposit x amount on your account
func (a Account) Deposit(amount int) {
	a.balance += amount
}

//print balance
func (a Account) Balance() int {
	return a.balance
}
