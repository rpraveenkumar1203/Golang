package prices

import (
	"fmt"

	"example.com/app.go/utils"
)

type Taxwithprice struct {
	Prices     []float64
	Taxrate    float64
	Totalprice map[string]string
}

func (t *Taxwithprice) LoadData() {

	data, dataerror := utils.Readfile("prices.txt")

	if dataerror != nil {
		fmt.Println(dataerror)
	}

	prices, conversionerror := utils.StringstoFloat(data)

	if conversionerror != nil {
		fmt.Println(conversionerror)
	}

	t.Prices = prices

}

func (t *Taxwithprice) PriceafterTax() {
	t.LoadData()

	PriceafterTax := make(map[string]string)

	for _, price := range t.Prices {
		taxincludePrice := price * (1 + t.Taxrate)
		PriceafterTax[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxincludePrice)

	}
	t.Totalprice = PriceafterTax

	// utils.Readfile(fmt.Sprintf("result_%.0f.json", t.taxrate*100))
	utils.WriteFile(fmt.Sprintf("result%.0f.json", t.Taxrate*100), t)

}

func TaxwithPrice(taxrate float64) *Taxwithprice {

	return &Taxwithprice{
		Prices:     []float64{},
		Taxrate:    taxrate,
		Totalprice: make(map[string]string),
	}
}
