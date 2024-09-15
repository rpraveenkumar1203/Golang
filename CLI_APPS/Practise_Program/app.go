package main

import (
	"example.com/app.go/prices"
)

func main() {

	taxrates := []float64{0, 0.01, 0.02, 0.03}

	for _, taxrate := range taxrates {
		pricewithtax := prices.TaxwithPrice(taxrate)
		pricewithtax.PriceafterTax()
	}
}
