package main

import (
	"fmt"
	"time"
)

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var userinput string
	fmt.Scan(&userinput)
	return userinput
}

type userdetails struct {
	fName string
	lName string
	dob   string
	time  time.Time
}

func ReturnUserDetails(d *userdetails) {
	fmt.Println(d.fName, d.lName, d.dob, d.time)
}

func main() {
	firstname := GetUserData("Please enter your first Name :- ")
	lastname := GetUserData("Please ENter your Late name :- ")
	dateofBirth := GetUserData("Please enter user DOB :- ")

	data := userdetails{
		fName: firstname,
		lName: lastname,
		dob:   dateofBirth,
		time:  time.Now(),
	}

	ReturnUserDetails(&data)

}

//output
//(base) pkr@rpraveenkumar:go run Structs.go
// Please enter your first Name :- PraveenKumar
// Please ENter your Late name :- Ramesh
// Please enter user DOB :- 12-03-1999
// PraveenKumar Ramesh 12-03-1999 2024-09-14 12:21:11.069066891 +0530 IST m=+14.154328567
