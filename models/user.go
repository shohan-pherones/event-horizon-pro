package models

import (
	"errors"

	"github.com/shohan-pherones/event-horizon-pro/db"
	"github.com/shohan-pherones/event-horizon-pro/utils"
)

type User struct {
	ID       int64
	Name     string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(name, email, password) VALUES(?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	isPasswordValid := utils.ComparingPasswordHash(u.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("credentials invalid")
	}

	return nil
}
