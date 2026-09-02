package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func setBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func gerBalance() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

var accountBalance float64 = gerBalance()

func main() {

	for {
		var depositMoney float64 = 0
		var withdrawAmount float64 = 0
		fmt.Println("Welcom to Go Bank")
		fmt.Println("What do you want to do")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposite money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is: ", accountBalance)
		case 2:
			fmt.Printf("How much amount you want to deposite: ")
			fmt.Scan(&depositMoney)
			if depositMoney <= 0 {
				fmt.Println("You can't deposit amount of ", depositMoney)
				continue
			} else {
				accountBalance += depositMoney
			}
			setBalance(accountBalance)
			fmt.Println("You have sucessfully deposite ", depositMoney, " and your account balance is ", accountBalance)

		case 3:
			fmt.Printf("How much amount you want to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw amount of ", withdrawAmount, "as Your accoun balance have only ", accountBalance)
				continue
			} else {
				accountBalance -= withdrawAmount
			}
			setBalance(accountBalance)
			fmt.Println("You have sucessfully withdraw ", withdrawAmount, " and your account balance is ", accountBalance)
		case 4:
			fmt.Println("Thank you for your visit")
			return
		default:
			fmt.Println("Invalid Selection")
		}
	}
}
