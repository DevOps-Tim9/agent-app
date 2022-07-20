package model

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type User struct {
	ID        int    `json:"id"`
	Auth0ID   string `json:"auth0_id"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email" gorm:"unique"`
	Password  string `json:"password"`
	Username  string `json:"user_name" gorm:"unique" validate:"required"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}
