package prices

import (
	"fmt"

	"example.com/app.go/utils"
)

type Taxwithprice struct {
	IoManager  utils.FileManager `json:"-"`
	Prices     []float64         `json:"Price_rate"`
	Taxrate    float64           `json:"tax_rate"`
	Totalprice map[string]string `json:"total_price"`
}

func (t *Taxwithprice) LoadData() {

	data, dataerror := t.IoManager.Readfile()

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

	t.IoManager.WriteFile(t)

}

func TaxwithPrice(FM utils.FileManager, taxrate float64) *Taxwithprice {

	return &Taxwithprice{
		IoManager:  FM,
		Prices:     []float64{},
		Taxrate:    taxrate,
		Totalprice: make(map[string]string),
	}
}
