package notes

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf
var scanln = fmt.Scanln
var scan = fmt.Scan

type userdata struct {
	Title   string
	Content string
	Time    time.Time
}

func (u userdata) DisplayNote() {

	printf("TITLE :-\n%v\nContent :-\n%v\n\n", u.Title, u.Content)

}

func (u userdata) Dateinv() error {

	fileName := strings.ReplaceAll(u.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0777)

}
func StoreNote(title, content string) (userdata, error) {

	if title == "" || content == "" {
		return userdata{}, fmt.Errorf("no data")
	}

	return userdata{

		Title:   title,
		Content: content,
		Time:    time.Now(),
	}, nil
}
