package banking

import "errors"

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

// Deposit : x amount on your account
func (a *Account) Deposit(amount int) {
	// (a Account) is receiver, make copy
	// (a *Account) is receiver, get origin
	a.balance += amount
}

// Balance : your account
func (a Account) Balance() int {
	return a.balance
}

var errNotEnoughMoney = errors.New("Can't withdraw your are poor")

// Withdraw : from your account
func (a *Account) Withdraw(amount int) error {
	// go do not have Exception, try catch
	// so, you have to return error, caller have to handle the error
	if a.balance < amount {
		return errNotEnoughMoney
	}

	a.balance -= amount
	return nil // nil is same as null in Java
}
