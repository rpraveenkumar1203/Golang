package main

import (
	"fmt"

	"example.com/app.go/prices"
	"example.com/app.go/utils"
)

func main() {

	taxrates := []float64{0.01, 0.02, 0.03}
	Cprocess := make([]chan bool, len(taxrates))
	Cerror := make([]chan error, len(taxrates))

	for index, taxrate := range taxrates {

		Cprocess[index] = make(chan bool)
		Cerror[index] = make(chan error)

		filemaanger := utils.New("prices.txt", fmt.Sprintf("result%.0f.json", taxrate*100))
		pricewithtax := prices.TaxwithPrice(filemaanger, taxrate)

		//climaanager := climanager.New()
		//pricewithtax := prices.TaxwithPrice(climaanager, taxrate)

		go pricewithtax.PriceafterTax(Cprocess[index], Cerror[index])

	}

	for index, _ := range taxrates {
		select {
		case err := <-Cerror[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-Cprocess[index]:
			fmt.Println("job done")
		}
	}
}
