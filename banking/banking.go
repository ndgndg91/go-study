package banking

// Account : public struct
// start Upper case is Public, start lower case is private
// if public, can create in other file
// if private, can create in only this file
type Account struct { //public struct
	owner   string //private field
	balance int    //private field
}

// NewAccount :
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
