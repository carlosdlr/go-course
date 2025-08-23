package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/util"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) SaveUser() error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := util.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // Optionally retrieve the last inserted ID
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() (bool, error) {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string
	if err := row.Scan(&u.ID, &hashedPassword); err != nil {
		return false, errors.New("invalid credentials")
	}

	err := util.CompareHashAndPassword(hashedPassword, u.Password)
	if err != nil {
		return false, errors.New("invalid credentials")
	}

	return true, nil
}
