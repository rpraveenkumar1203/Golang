package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatValue(filename string) (float64, error) {

	file, err := os.ReadFile(filename)

	if err != nil {
		return 1110, errors.New("no file found")
	}
	balanceData := string(file)
	balanceAmount, err := strconv.ParseFloat(balanceData, 64)

	if err != nil {
		return 1110, errors.New("unable to parse the balance ")
	}
	return balanceAmount, nil

}

func WriteFloatValue(value float64, filename string) {

	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0777)

}
