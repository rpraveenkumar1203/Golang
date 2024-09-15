package prices

import "fmt"

type taxwithprice struct {
	prices     []float64
	taxrate    float64
	totalprice map[string]float64
}

func (t taxwithprice) PriceafterTax() {

	priceaftertax := make(map[string]float64)

	for _, price := range t.prices {
		priceaftertax[fmt.Sprintf("%.2f", price)] = price * (1 * t.taxrate)
	}
	fmt.Println(priceaftertax)

}

func TaxwithPrice(taxrate float64) *taxwithprice {

	return &taxwithprice{
		prices:  []float64{10, 20, 30},
		taxrate: taxrate,
	}
}
