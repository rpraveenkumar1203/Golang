package utils

import (
	"errors"
	"strconv"
)

/*
GETS INPUT AS SLICE OF STRINGS AND CONVERT ALL OF THEM TO SLICE OF FLOATS
*/

func StringstoFloat(strings []string) ([]float64, error) {

	var floatvalues []float64

	for _, string := range strings {
		floatvalue, err := strconv.ParseFloat(string, 64)
		if err != nil {
			return nil, errors.New("float conversion fdalied")
		}
		floatvalues = append(floatvalues, floatvalue)
	}
	return floatvalues, nil
}
