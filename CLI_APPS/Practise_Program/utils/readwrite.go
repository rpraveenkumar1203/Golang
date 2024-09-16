package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func Readfile(path string) ([]string, error) {

	file, err := os.Open(path)
	readfile := bufio.NewScanner(file)
	defer file.Close()
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

func WriteFile(path string, data any) error {

	createdfile, createfileerr := os.Create(path)
	if createfileerr != nil {
		return errors.New("unable to create file ")
	}
	defer createdfile.Close()
	dataencode := json.NewEncoder(createdfile)

	dataencodeError := dataencode.Encode(data)

	if dataencodeError != nil {
		return errors.New("failed to write file ")

	}
	return nil

}
