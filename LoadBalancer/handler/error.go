package handler

import (
	"fmt"
	"os"
)

func Error(err error) {
	if err != nil {
		fmt.Printf("ERROR :- %v \n", err)
		os.Exit(1)
	}
}
