package main

import "fmt"

func main() {

	// myLoop:
	for {
		var accountBalance float64 = 1000
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

		wantsCheckBalance := choice == 1
		wantDepositeMoney := choice == 2
		wantsWithdrawBalance := choice == 3
		wantexit := choice == 4

		if wantsCheckBalance {
			fmt.Println("Your account balance is: ", accountBalance)
		} else if wantDepositeMoney {
			fmt.Printf("How much amount you want to deposite: ")
			fmt.Scan(&depositMoney)
			if depositMoney <= 0 {
				fmt.Println("You can't deposit amount of ", depositMoney)
				continue
			} else {
				accountBalance += depositMoney
			}

			fmt.Println("You have sucessfully deposite ", depositMoney, " and your account balance is ", accountBalance)
		} else if wantsWithdrawBalance {
			fmt.Printf("How much amount you want to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw amount of ", withdrawAmount, "as Your accoun balance have only ", accountBalance)
				continue
			} else {
				accountBalance -= withdrawAmount
			}
			fmt.Println("You have sucessfully withdraw ", withdrawAmount, " and your account balance is ", accountBalance)

		} else if wantexit {
			fmt.Print("Thank you for your visit")
			break
		} else {
			fmt.Print("Invalid Selection")
		}
	}
}
