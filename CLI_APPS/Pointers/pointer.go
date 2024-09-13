package main

func main() {

	age := 32
	ageP := &age

	getadultage(age)

}

func getadultage(age *int) int {

	return ageP - 18

}
