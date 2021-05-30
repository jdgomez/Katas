package main

import (
	"fmt"
	"testing"
)

func TestNewAccount(t *testing.T) {

	account := account{}
	if account.balance != 0 {
		t.Errorf("Error: new accounts doesn't start empty")
	}
}

func TestDeposit(t *testing.T) {

	depositAmount := 100
	account := account{}
	originalBalance := account.getBalance()

	account.deposit(depositAmount)

	if account.getBalance() != (originalBalance + depositAmount) {
		t.Errorf("Error: the balance didn't increase as expected")
	}
}

func TestWithdraw(t *testing.T) {

	withdrawAmount := 55
	account := account{}
	originalBalance := account.getBalance()

	account.withdraw(withdrawAmount)

	if account.getBalance() != (originalBalance - withdrawAmount) {
		fmt.Println("Account:", account.getBalance())
		t.Errorf("Error: the balance didn't decrease as expected")
	}
}

func TestPrintStatement(t *testing.T) {

	expectedString := "Date\t\tAmount\tBalance\n30.05.2021\t+500\t500\n30.05.2021\t-100\t400\n"
	depositAmount := 500
	withdrawAmount := 100
	account := account{}

	account.deposit(depositAmount)
	account.withdraw(withdrawAmount)

	if account.printStatement() != expectedString {
		fmt.Println("Expected:\n", expectedString)
		fmt.Println("Account:\n", account.printStatement())
		t.Errorf("Error: the string isn't right")
	}
}
