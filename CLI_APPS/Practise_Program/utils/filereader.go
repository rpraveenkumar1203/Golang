package utils

import (
	"bufio"
	"errors"
	"os"
)

func Readfile(path string) ([]string, error) {

	file, err := os.Open(path)
	readfile := bufio.NewScanner(file)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	var value []string

	for readfile.Scan() {
		value = append(value, readfile.Text())

	}
	err = readfile.Err()

	if err != nil {
		return nil, errors.New("reading error")
	}

	return value, nil

}
