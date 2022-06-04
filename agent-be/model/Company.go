package model

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type Company struct {
	ID           int    `json:"id"`
	OwnerAuth0ID string `json:"owner_id"`
	Name         string `json:"name" validate:"required"`
	Contact      string `json:"contact" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Approved     bool   `json:"approved"`
}

func (c *Company) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *Company) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

func (c *Company) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
