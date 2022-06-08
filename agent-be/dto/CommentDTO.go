package dto

import "time"

type CommentDTO struct {
	ID           int
	UserOwnerID  int
	Salary       float32
	Position     string
	Description  string
	CreationDate time.Time
	CompanyID    int
}
