package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// GETS FILEPATH AS .TXT AND RETURNS VALUES AS SLICE OF STRINGS

func (FM FileManager) Readfile() ([]string, error) {

	file, err := os.Open(FM.InputFilePath)
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

// GETS FILE_PATH AND UPDATE THE DATA AS JSON

func (FM FileManager) WriteFile(data any) error {

	createdfile, createfileerr := os.Create(FM.OutputFilePath)
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

func New(inputfile, outpufile string) FileManager {

	return FileManager{
		InputFilePath:  inputfile,
		OutputFilePath: outpufile,
	}
}
