package mapper

import (
	"agent-app/src/dto"
	"agent-app/src/model"
)

func RegistrationRequestDTOToUser(registeredUserDto *dto.RegistrationRequestDTO) *model.User {
	var user model.User

	user.Username = registeredUserDto.Username
	user.FirstName = registeredUserDto.FirstName
	user.LastName = registeredUserDto.LastName
	user.Email = registeredUserDto.Email
	user.Password = registeredUserDto.Password

	return &user
}

func UserToDTO(userEntity *model.User) *dto.UserResponseDTO {
	var user dto.UserResponseDTO

	user.ID = userEntity.ID
	user.Auth0ID = userEntity.Auth0ID
	user.FirstName = userEntity.FirstName
	user.LastName = userEntity.LastName
	user.Email = userEntity.Email
	user.Username = userEntity.Username

	return &user
}
