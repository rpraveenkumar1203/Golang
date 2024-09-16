package prices

import (
	"bufio"
	"fmt"
	"os"

	conversion "example.com/app.go/utils"
)

type taxwithprice struct {
	prices     []float64
	taxrate    float64
	totalprice map[string]float64
}

func (t *taxwithprice) LoadData() {

	dataFile, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println(err)
	}
	datascanned := bufio.NewScanner(dataFile)

	var ScannedData []string

	for datascanned.Scan() {
		ScannedData = append(ScannedData, datascanned.Text())
	}
	err = datascanned.Err()
	if err != nil {
		fmt.Println(err)
		dataFile.Close()
		return
	}

	prices, _ := conversion.StringstoFloat(ScannedData)

	t.prices = prices

}

func (t taxwithprice) PriceafterTax() {
	t.LoadData()

	priceaftertax := make(map[string]float64)

	for _, price := range t.prices {
		priceaftertax[fmt.Sprintf("%.2f", price)] = price * (1 + t.taxrate)
	}
	fmt.Println(priceaftertax)

}

func TaxwithPrice(taxrate float64) *taxwithprice {

	return &taxwithprice{
		prices:  []float64{10, 20, 30},
		taxrate: taxrate,
	}
}
