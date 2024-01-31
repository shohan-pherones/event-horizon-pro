package models

import "github.com/shohan-pherones/event-horizon-pro/db"

type User struct {
	ID       int64
	Name     string `binding:"required"`
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

	result, err := stmt.Exec(u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}
