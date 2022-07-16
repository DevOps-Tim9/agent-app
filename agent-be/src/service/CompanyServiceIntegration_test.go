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

type CompanyServiceIntegrationTestSuite struct {
	suite.Suite
	service   CompanyService
	db        *gorm.DB
	companies []model.Company
	users     []model.User
}

func (suite *CompanyServiceIntegrationTestSuite) SetupSuite() {
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

	db.AutoMigrate(model.Company{})
	db.AutoMigrate(model.User{})

	db.Where("1=1").Delete(model.Company{})
	db.Where("1=1").Delete(model.User{})

	auth0Client := auth0.NewAuth0Client(os.Getenv("AUTH0_DOMAIN"), os.Getenv("AUTH0_CLIENT_ID"), os.Getenv("AUTH0_CLIENT_SECRET"), os.Getenv("AUTH0_AUDIENCE"))

	repo := repository.CompanyRepository{Database: db}
	userRepo := repository.UserRepository{Database: db}

	suite.db = db

	suite.service = CompanyService{
		CompanyRepo: &repo,
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

	suite.companies = []model.Company{
		{
			Name:         "Test",
			Description:  "Test desc",
			Contact:      "Test contact",
			Approved:     true,
			OwnerAuth0ID: "auth0123",
		},
		{
			Name:         "Test 1",
			Description:  "Test desc",
			Contact:      "Test contact",
			Approved:     false,
			OwnerAuth0ID: "auth0123",
		},
	}

	tx = suite.db.Begin()

	tx.Create(&suite.companies[0])
	tx.Create(&suite.companies[1])

	tx.Commit()
}

func TestCompanyServiceIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(CompanyServiceIntegrationTestSuite))
}

func (suite *CompanyServiceIntegrationTestSuite) TestIntegrationCompanyService_GetAll_ApprovedExist() {
	companies, err := suite.service.GetAll(1)

	assert.NotNil(suite.T(), companies)
	assert.Equal(suite.T(), 1, len(companies))
	assert.Nil(suite.T(), err)
}

func (suite *CompanyServiceIntegrationTestSuite) TestIntegrationCompanyService_GetAll_NotApprovedExist() {
	companies, err := suite.service.GetAll(0)

	assert.NotNil(suite.T(), companies)
	assert.Equal(suite.T(), 1, len(companies))
	assert.Nil(suite.T(), err)
}

func (suite *CompanyServiceIntegrationTestSuite) TestIntegrationCompanyService_Approve_CompanyExists() {
	request := dto.ApproveCompanyDTO{
		ID:      1,
		Approve: true,
	}

	err := suite.service.Approve(&request)
	assert.Nil(suite.T(), err)
}

func (suite *CompanyServiceIntegrationTestSuite) TestIntegrationCompanyService_Approve_CompanyDoesNotExist() {
	request := dto.ApproveCompanyDTO{
		ID:      10000,
		Approve: true,
	}

	suite.service.Approve(&request)
	assert.True(suite.T(), true)
}

func (suite *CompanyServiceIntegrationTestSuite) TestIntegrationCompanyService_Register_Pass() {
	companyDTO := dto.CompanyRequestDTO{
		Name:        "Test",
		Description: "Test description",
		Contact:     "Test contact",
	}
	auth0ID := "123auth0"

	id, err := suite.service.Register(&companyDTO, auth0ID)
	assert.Nil(suite.T(), err)
	assert.NotEqual(suite.T(), -1, id)
}

func (suite *CompanyServiceIntegrationTestSuite) TestIntegrationCompanyService_Register_Fail() {
	companyDTO := dto.CompanyRequestDTO{
		Description: "Test description",
		Contact:     "Test contact",
	}
	auth0ID := "123auth0"

	id, err := suite.service.Register(&companyDTO, auth0ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), -1, id)
}

func (suite *CompanyServiceIntegrationTestSuite) TestIntegrationCompanyService_Update_CompanyDoesNotExist() {
	company := dto.CompanyUpdateDTO{
		ID:          10000000000000,
		Name:        "new name",
		Contact:     "new contact",
		Description: "new description",
	}
	auth0ID := "123auth0"

	response, err := suite.service.Update(&company, auth0ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), nil, response)
}

func (suite *CompanyServiceIntegrationTestSuite) TestIntegrationCompanyService_Update_OwnerNotUpdatingCompany() {
	company := dto.CompanyUpdateDTO{
		ID:          1,
		Name:        "new name",
		Contact:     "new contact",
		Description: "new description",
	}
	auth0ID := "auth0124"

	response, err := suite.service.Update(&company, auth0ID)
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), response)
}

func (suite *CompanyServiceIntegrationTestSuite) TestIntegrationCompanyService_Update_Pass() {
	company := dto.CompanyUpdateDTO{
		ID:          1,
		Name:        "new name",
		Contact:     "new contact",
		Description: "new description",
	}
	auth0ID := "auth0123"

	response, err := suite.service.Update(&company, auth0ID)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), response)
	assert.Equal(suite.T(), company.Name, response.Name)
	assert.Equal(suite.T(), company.Contact, response.Contact)
	assert.Equal(suite.T(), company.Description, response.Description)
}
