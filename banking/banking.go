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
