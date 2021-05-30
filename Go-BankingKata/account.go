package main

import (
	"strconv"
	"time"
)

type operation struct {
	date      time.Time
	amount    int
	isDeposit bool
}
type account struct {
	balance int
	history []operation
}

func (a account) getBalance() int {
	for _, o := range a.history {
		a.applyOperation(o)
	}
	return a.balance
}

func (a *account) deposit(amount int) {
	a.history = append(a.history, operation{
		time.Now(),
		amount,
		true,
	})
}

func (a *account) withdraw(amount int) {
	a.history = append(a.history, operation{
		time.Now(),
		amount,
		false,
	})
}

func (a account) printStatement() string {
	statement := "Date\t\tAmount\tBalance\n"
	for _, o := range a.history {
		statement = statement + o.date.Format("02.01.2006") + "\t"
		statement = statement + o.getSign() + strconv.Itoa(o.amount) + "\t"
		a.applyOperation(o)
		statement = statement + strconv.Itoa(a.balance) + "\n"
	}

	return statement
}

func (o operation) getSign() string {
	if o.isDeposit {
		return "+"
	}
	return "-"
}

func (a *account) applyOperation(o operation) {
	if o.isDeposit {
		a.balance += o.amount
	} else {
		a.balance -= o.amount
	}
}