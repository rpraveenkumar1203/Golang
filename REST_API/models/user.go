package models

import (
	"errors"

	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/db"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/utils"
)

type Userdata struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *Userdata) Save() error {

	query := "INSERT INTO USERS (email,password) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	hashed_password, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashed_password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id

	return err

}

func (u *Userdata) Validatelogin() error {

	query := "SELECT password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var password string

	err := row.Scan(&password)

	if err != nil {
		return errors.New("invalid data ")
	}

	validPassword := utils.CheckHashPassword(u.Password, password)

	if !validPassword {

		return errors.New("invalid data ")
	}

	return nil

}
