package main

import (
	"fmt"
	"math" 
)

var print = fmt.Print
var println = fmt.Println

func main() {
	println("Building investment Calculator APP")

	// input

	var investmentAmount float64 = 100000
	var returnRate float64 = 5.5
	var tenure float64 = 10

	// output
	var expectedReturnAmount = investmentAmount * math.Pow(1+returnRate/100, tenure)

	println(expectedReturnAmount)
}
