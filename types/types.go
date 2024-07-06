package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user User) error
}

type User struct {
	ID        int       `json:"_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	UserName  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}
