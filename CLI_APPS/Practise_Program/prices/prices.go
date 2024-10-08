package prices

import (
	"fmt"

	iomanager "example.com/app.go/ioManager"
	"example.com/app.go/utils"
)

type Taxwithprice struct {
	IoManager  iomanager.IOmanager `json:"-"`
	Prices     []float64           `json:"Price_rate"`
	Taxrate    float64             `json:"tax_rate"`
	Totalprice map[string]string   `json:"total_price"`
}

func (t *Taxwithprice) LoadData() error {

	data, dataerror := t.IoManager.Readfile()

	if dataerror != nil {
		fmt.Println(dataerror)
	}

	prices, conversionerror := utils.StringstoFloat(data)

	if conversionerror != nil {
		fmt.Println(conversionerror)
	}

	t.Prices = prices
	return nil

}

func (t *Taxwithprice) PriceafterTax(Cprocess chan bool, Cerror chan error) {
	err := t.LoadData()

	if err != nil {
		Cerror <- err
		return

	}

	PriceafterTax := make(map[string]string)

	for _, price := range t.Prices {
		taxincludePrice := price * (1 + t.Taxrate)
		PriceafterTax[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxincludePrice)

	}
	t.Totalprice = PriceafterTax

	t.IoManager.WriteFile(t)
	Cprocess <- true

}

func TaxwithPrice(IO iomanager.IOmanager, taxrate float64) *Taxwithprice {

	return &Taxwithprice{
		IoManager:  IO,
		Prices:     []float64{},
		Taxrate:    taxrate,
		Totalprice: make(map[string]string),
	}
}
