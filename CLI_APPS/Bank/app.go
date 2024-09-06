package main

import "fmt"

func main() {

	var userAmount float64 = 1000

	fmt.Println("Please select from below options :")

	for {

		fmt.Println("1. Deposit Money")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")

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
				fmt.Printf("The total amount after withdrawl is %.2f \n", userAmount)
			} else {
				fmt.Printf("Please enter between 50 dollars and  %.2f dollars \n", userAmount)

			}

		} else if userChoice == 3 {

			fmt.Printf("Your current balance is %.2f \n", userAmount)
			break

		} else {
			fmt.Println("Logged out of session successfully")
			break
		}

	}
	fmt.Println("Thanks for using our service")

}
