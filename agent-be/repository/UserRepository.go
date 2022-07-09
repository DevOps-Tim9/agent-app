package repository

import (
	"agent-app/dto"
	"agent-app/mapper"
	"agent-app/model"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type IUserRepository interface {
	AddUser(*model.User) (int, error)
	DeleteUser(int) error
	GetByID(int) (*model.User, error)
	GetByAuth0ID(string) (*model.User, error)
	GetByEmail(string) (*dto.UserResponseDTO, error)
	Update(*model.User) (*dto.UserResponseDTO, error)
	CreateAdmin([]model.User)
}

func NewUserRepository(database *gorm.DB) IUserRepository {
	return &UserRepository{
		database,
	}
}

type UserRepository struct {
	Database *gorm.DB
}

func (repo *UserRepository) CreateAdmin(admins []model.User) {
	for i := 0; i < len(admins); i++ {
		if repo.Database.Where("email = ?", admins[i].Email).RowsAffected == 0 {
			repo.AddUser(&admins[i])
		}
	}
}

func (repo *UserRepository) AddUser(user *model.User) (int, error) {
	result := repo.Database.Create(user)

	if result.Error != nil {
		return -1, result.Error
	}

	return user.ID, nil
}

func (repo *UserRepository) DeleteUser(id int) error {
	result := repo.Database.Delete(&model.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *UserRepository) GetByID(id int) (*model.User, error) {
	userEntity := model.User{
		ID: id,
	}
	if err := repo.Database.Where("ID = ?", id).First(&userEntity).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("User with ID %d not found", id))
	}

	return &userEntity, nil
}

func (repo *UserRepository) GetByAuth0ID(id string) (*model.User, error) {
	userEntity := model.User{
		Auth0ID: id,
	}
	if err := repo.Database.Where("auth0_id = ?", id).First(&userEntity).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("User with auth0ID %s not found", id))
	}

	return &userEntity, nil
}

func (repo *UserRepository) GetByEmail(email string) (*dto.UserResponseDTO, error) {
	userEntity := model.User{
		Email: email,
	}
	if err := repo.Database.Where("email = ?", email).First(&userEntity).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("User with email %s not found", email))
	}

	return mapper.UserToDTO(&userEntity), nil
}

func (repo *UserRepository) Update(user *model.User) (*dto.UserResponseDTO, error) {
	result := repo.Database.Save(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return mapper.UserToDTO(user), nil
}
