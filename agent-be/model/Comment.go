package model

import (
	"encoding/json"
	"io"
	"time"

	"github.com/go-playground/validator"
)

type Comment struct {
	ID               int       `json:"id"`
	UserOwnerID      int       `json:"user_owner_id" gorm:"TYPE:integer REFERENCES users"`
	Salary           float32   `json:"salary" validate:"required"`
	Position         string    `json:"position" validate:"required"`
	Rating           string    `json:"rating" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	InterviewProcess string    `json:"interview_process"`
	CreationDate     time.Time `json:"creation_date" validate:"required" validate:"required"`
	CompanyID        int       `json:"company_id" gorm:"TYPE:integer REFERENCES companies" validate:"required"`
}

func (c *Comment) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *Comment) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

func (c *Comment) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
