package service

import (
	"agent-app/src/auth0"
	"agent-app/src/dto"
	"agent-app/src/model"
	"agent-app/src/repository"
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceIntegrationTestSuite struct {
	suite.Suite
	service UserService
	db      *gorm.DB
	users   []model.User
}

func (suite *UserServiceIntegrationTestSuite) SetupSuite() {
	host := os.Getenv("DATABASE_DOMAIN")
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_SCHEMA")
	port := os.Getenv("DATABASE_PORT")

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		name,
		port,
	)
	db, _ := gorm.Open("postgres", connectionString)

	db.AutoMigrate(model.User{})

	db.Where("1=1").Delete(model.User{})

	auth0Client := auth0.NewAuth0Client(os.Getenv("AUTH0_DOMAIN"), os.Getenv("AUTH0_CLIENT_ID"), os.Getenv("AUTH0_CLIENT_SECRET"), os.Getenv("AUTH0_AUDIENCE"))
	userRepo := repository.UserRepository{Database: db}

	suite.db = db

	suite.service = UserService{
		Auth0Client: auth0Client,
		UserRepo:    &userRepo,
	}

	suite.users = []model.User{
		{
			FirstName: "Name",
			LastName:  "Surname",
			Email:     "test@test.com",
			Password:  "password123",
			Username:  "username",
			Auth0ID:   "auth0123",
		},
	}

	tx := suite.db.Begin()

	tx.Create(&suite.users[0])

	tx.Commit()
}

func TestUserServiceIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceIntegrationTestSuite))
}

func (suite *UserServiceIntegrationTestSuite) TestIntegrationUserService_GetByEmail_UserExists() {
	user, err := suite.service.GetByEmail("test@test.com")

	assert.NotNil(suite.T(), user)
	assert.Nil(suite.T(), err)
}

func (suite *UserServiceIntegrationTestSuite) TestIntegrationUserService_GetByEmail_UserDoesNotExist() {
	user, err := suite.service.GetByEmail("testemail123@test.com")

	assert.Nil(suite.T(), user)
	assert.NotNil(suite.T(), err)
}

func (suite *UserServiceIntegrationTestSuite) TestIntegrationUserService_Register_Password6CharsLong() {
	userDTO := dto.RegistrationRequestDTO{
		FirstName: "Name",
		LastName:  "Surname",
		Email:     "test@test.com",
		Password:  "pas123",
		Username:  "username",
	}

	id, err := suite.service.Register(&userDTO)

	assert.Equal(suite.T(), -1, id)
	assert.NotNil(suite.T(), err)
}

func (suite *UserServiceIntegrationTestSuite) TestIntegrationUserService_Register_PasswordDoesNotContainNumber() {
	userDTO := dto.RegistrationRequestDTO{
		FirstName: "Name",
		LastName:  "Surname",
		Email:     "test@test.com",
		Password:  "passwords",
		Username:  "username",
	}

	id, err := suite.service.Register(&userDTO)

	assert.Equal(suite.T(), -1, id)
	assert.NotNil(suite.T(), err)
}

func (suite *UserServiceIntegrationTestSuite) TestIntegrationUserService_Register_SameEmail() {
	userDTO := dto.RegistrationRequestDTO{
		FirstName: "Name",
		LastName:  "Surname",
		Email:     "test@test.com",
		Password:  "passwords",
		Username:  "usernameNew",
	}

	id, err := suite.service.Register(&userDTO)

	assert.Equal(suite.T(), -1, id)
	assert.NotNil(suite.T(), err)
}

func (suite *UserServiceIntegrationTestSuite) TestIntegrationUserService_Register_SameUsername() {
	userDTO := dto.RegistrationRequestDTO{
		FirstName: "Name",
		LastName:  "Surname",
		Email:     "testemail@test.com",
		Password:  "passwords",
		Username:  "username",
	}

	id, err := suite.service.Register(&userDTO)

	assert.Equal(suite.T(), -1, id)
	assert.NotNil(suite.T(), err)
}
