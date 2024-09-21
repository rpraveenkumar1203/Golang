package models

import "github.com/rpraveenkumar1203/Golang/tree/main/REST_API/db"

type Userdata struct {
	ID       int64
	Email    string
	Password string
}

func (u *Userdata) Save() error {

	query := "INSERT INTO USERS (email,password) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id

	return err

}
