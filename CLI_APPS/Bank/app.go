package main

import (
	"fmt"

	fileops "example.com/bank/ops"
	welcometemplate "example.com/bank/template"
)

const balanceFile = "balance.txt"

func main() {

	//var userAmount float64 = 1000
	var userAmount, err = fileops.GetFloatValue(balanceFile)

	if err != nil {
		fmt.Println("---------------Database not found --------")
		panic("No data found")
	}

	fmt.Println("Please select from below options :")

	for {

		welcometemplate.WelcomeTemplate()

		var userChoice int

		fmt.Println("Please enter a option ")
		fmt.Scan(&userChoice)

		//fmt.Println(userChoice)

		if userChoice == 1 {

			fmt.Println("Please Enter amount you are going to deposit")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if !(depositAmount <= 100) {
				userAmount = userAmount + depositAmount
				fileops.WriteFloatValue(userAmount, balanceFile)

				fmt.Printf("The total amount after deposit is %.2f \n", userAmount)

			} else {
				fmt.Println("Please enter a amount above 100 dollars ")
			}

		} else if userChoice == 2 {

			fmt.Println("Please enter the amount you're going to withdraw :- ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if !(withdrawAmount > userAmount) && !(withdrawAmount <= 50) {
				userAmount = userAmount - withdrawAmount
				fileops.WriteFloatValue(userAmount, balanceFile)
				fmt.Printf("The total amount after withdrawl is %.2f \n", userAmount)
			} else {
				fmt.Printf("Please enter between 50 dollars and  %.2f dollars \n", userAmount)

			}

		} else if userChoice == 3 {

			fmt.Printf("Your current balance is %.2f \n", userAmount)
			continue

		} else if userChoice == 4 {
			fmt.Println("\nLogged out of session successfully")
			break
		} else {
			fmt.Println("Invalid input")
			continue
		}

	}
	fmt.Println("\nThanks for using our service")

}
