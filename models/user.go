package models

import (
	"net/http"

	"example.com/rest/db"
	"example.com/rest/utils"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
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
		"INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id",
		u.Email, u.Password,
	).Scan(&u.Id)
	if err != nil {
		return err, uuid.Nil
	}
	return nil, u.Id
}

func GetUserByEmail(r *http.Request, email string) (*User, error) {
	var user User
	err := db.Pool.QueryRow(r.Context(),
		"SELECT id, email, password_hash FROM users WHERE email = $1",
		email,
	).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u User) ValidateCredentials(r *http.Request, password string) bool {
	isValid := utils.CheckPasswordHash(u.Password, password)
	return isValid
}
