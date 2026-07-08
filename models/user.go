package models

import (
	"net/http"

	"example.com/rest/db"
	"example.com/rest/utils"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

func (u User) Save(r *http.Request) (error, uuid.UUID) {

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err, uuid.Nil
	}
	u.Password = hashedPassword

	err = db.Pool.QueryRow(r.Context(),
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id",
		u.Email, u.Password,
	).Scan(&u.Id)
	if err != nil {
		return err, uuid.Nil
	}
	return nil, u.Id
}
