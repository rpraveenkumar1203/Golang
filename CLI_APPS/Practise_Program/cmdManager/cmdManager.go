package climanager

import (
	"fmt"
)

type Clidata struct {
}

func (CLI Clidata) Readfile() ([]string, error) {

	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (CLI Clidata) WriteFile(data any) error {
	fmt.Println(data)
	return nil
}

func New() Clidata {
	return Clidata{}
}
