package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, args := range os.Args[1:] {
		sep = " "
		s = s + sep + args
		fmt.Println(s)
	}
}

// (base) pkr@rpraveenkumar:/media/pkr/bala/pygodev/Golang/Basic_programs$ go run sample2.go 1 2 3 4 5
// 1
// 1 2
// 1 2 3
// 1 2 3 4
// 1 2 3 4 5
// (base) pkr@rpraveenkumar:/media/pkr/bala/pygodev/Golang/Basic_programs$ go run sample2.go 1
// 1
// (base) pkr@rpraveenkumar:/media/pkr/bala/pygodev/Golang/Basic_programs$ go run sample2.go 1 3
// 1
// 1 3
// (base) pkr@rpraveenkumar:/media/pkr/bala/pygodev/Golang/Basic_programs$
