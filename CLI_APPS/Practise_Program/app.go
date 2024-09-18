package main

import (
	"fmt"

	"example.com/app.go/prices"
	"example.com/app.go/utils"
)

func main() {

	taxrates := []float64{0.01, 0.02, 0.03}

	for _, taxrate := range taxrates {

		filemaanger := utils.New("prices.txt", fmt.Sprintf("result%.0f.json", taxrate*100))
		pricewithtax := prices.TaxwithPrice(filemaanger, taxrate)

		//climaanager := climanager.New()
		//pricewithtax := prices.TaxwithPrice(climaanager, taxrate)

		pricewithtax.PriceafterTax()
	}
}
