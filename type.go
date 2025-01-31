package main

import (
	"math/rand"
	"strconv"
)

type Account struct {
	ID        string
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        strconv.Itoa(rand.Intn(10000)),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
	}
}
