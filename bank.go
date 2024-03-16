package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}
func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		return
		// panic("Can't continue")
	}
	fmt.Println("Welcome to go bank")
	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated, New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				// return
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated, New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Good bye!")
			fmt.Println("Thanks for choosing our bank")

			return
			// break
		}
		// if choice == 1 {
		// 	fmt.Println("Your balance is: ", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your Deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		// return
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated, New amount: ", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Your withdraw: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		// return
		// 		continue
		// 	}
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Invalid amount. You can't withdraw more than you have.")
		// 		// return
		// 		continue
		// 	}
		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Balance updated, New amount: ", accountBalance)
		// } else {
		// 	fmt.Println("Good bye!")
		// 	// return
		// 	break
		// }

	}
}
