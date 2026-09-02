package main

import (
	"example.com/bank/fileops"
	"fmt"
)

const accountBalanceFile = "balance.txt"

// func setBalance(balance float64) {
// 	balanceText := fmt.Sprint(balance)
// 	os.WriteFile("balance.txt", []byte(balanceText), 0644)
// }

// func getFloatFromFiles(fileName string) (float64, error) {
// 	data, err := os.ReadFile(fileName)

// 	if err != nil {
// 		return 1000, errors.New("Failed to parse value")
// 	}
// 	balanceText := string(data)
// 	balance, _ := strconv.ParseFloat(balanceText, 64)
// 	return balance, nil
// }

var accountBalance, _ = fileops.getFloatFromFiles(accountBalanceFile)

func main() {

	for {
		var depositMoney float64 = 0
		var withdrawAmount float64 = 0

		presentOptions()

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
