package dto

type UserResponseDTO struct {
	ID        int
	Auth0ID   string
	FirstName string
	LastName  string
	Email     string
	Username  string
}
