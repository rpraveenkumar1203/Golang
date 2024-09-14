package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/app/notes"
)

var print = fmt.Print
var println = fmt.Println

func checkError(err error, result string) {
	if err != nil {
		println(err, "::Fail")
	}
	println(result, "::Pass")
}

func Input(userprompt string) string {
	for {
		print(userprompt)
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			checkError(err, "")
			continue
		}
		text = strings.TrimSpace(text)
		if text == "" {
			println("Input cannot be empty. Please try again.")
			continue
		}
		return text
	}
}

func getuserdata() (string, string) {

	title := Input("Enter the title :- ")
	content := Input("Enter the content :-")
	return title, content
}

func main() {

	title, content := getuserdata()
	storedata, err := notes.StoreNote(title, content)
	checkError(err, "data got for stored data ")
	storedata.DisplayNote()
	checkError(storedata.Dateinv(), "Copying done")

}
