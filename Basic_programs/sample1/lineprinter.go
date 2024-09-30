package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		sep = " "
		s = s + sep + os.Args[i]
		fmt.Println(s)
	}
}

/*
Output

(base) pkr@rpraveenkumar:/media/pkr/bala/pygodev/Golang/Basic_programs$ go run sample1/lineprinter.go 1 2 3 4 5
 1
 1 2
 1 2 3
 1 2 3 4
 1 2 3 4 5
(base) pkr@rpraveenkumar:/media/pkr/bala/pygodev/Golang/Basic_programs$

*/
