package conversion

import (
	"errors"
	"strconv"
)

func StringstoFloat(strings []string) ([]float64, error) {

	var floatvalues []float64

	for _, string := range strings {
		floatvalue, err := strconv.ParseFloat(string, 64)
		if err != nil {
			return nil, errors.New("float conversaion fdalied")
		}
		floatvalues = append(floatvalues, floatvalue)
	}
	return floatvalues, nil
}
