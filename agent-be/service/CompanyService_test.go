package service

import (
	"agent-app/auth0"
	"agent-app/dto"
	"agent-app/mapper"
	"agent-app/model"
	"agent-app/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CompanyServiceUnitTestsSuite struct {
	suite.Suite
	companyRepositoryMock *repository.CompanyRepositoryMock
	auth0ClientMock       *auth0.Auth0ClientMock
	service               ICompanyService
	userRepositoryMock    *repository.UserRepositoryMock
}

func TestCompanyServiceUnitTestsSuite(t *testing.T) {
	suite.Run(t, new(CompanyServiceUnitTestsSuite))
}

func (suite *CompanyServiceUnitTestsSuite) SetupSuite() {
	suite.companyRepositoryMock = new(repository.CompanyRepositoryMock)
	suite.userRepositoryMock = new(repository.UserRepositoryMock)
	suite.auth0ClientMock = new(auth0.Auth0ClientMock)
	suite.service = NewCompanyService(suite.companyRepositoryMock, suite.userRepositoryMock, suite.auth0ClientMock)
}

func (suite *CompanyServiceUnitTestsSuite) TestNewCompanyService() {
	assert.NotNil(suite.T(), suite.service, "Service is nil")
}

func (suite *CompanyServiceUnitTestsSuite) TestCompanyService_Register_ValidDataProvided() {
	companyDTO := dto.CompanyRequestDTO{
		Name:        "Test",
		Description: "Test description",
		Contact:     "Test contact",
	}
	auth0ID := "123auth0"

	suite.companyRepositoryMock.On("AddCompany", mock.AnythingOfType("*model.Company")).Return(1, nil).Once()

	id, err := suite.service.Register(&companyDTO, auth0ID)

	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 1, id)
}

func (suite *CompanyServiceUnitTestsSuite) TestCompanyService_GetAll_GetsAllCompanies() {
	company := model.Company{
		ID:           1,
		Name:         "Test",
		Description:  "Test desc",
		Contact:      "Test contact",
		Approved:     true,
		OwnerAuth0ID: "auth0123",
	}
	companyDTO := dto.CompanyResponseDTO{
		ID:          1,
		Name:        "Test",
		Description: "Test desc",
		Contact:     "Test contact",
		Owner:       "First Last",
		OwnerId:     "auth0123",
	}
	user := model.User{
		ID:        1,
		Auth0ID:   "auth0123",
		FirstName: "First",
		LastName:  "Last",
		Email:     "email@email",
		Password:  "pass123",
		Username:  "Username",
	}
	var companies []*model.Company
	companies = append(companies, &company)

	var companiesDTO []*dto.CompanyResponseDTO
	companiesDTO = append(companiesDTO, &companyDTO)

	suite.companyRepositoryMock.On("GetAll", 1).Return(companies, nil).Once()
	suite.userRepositoryMock.On("GetByAuth0ID", "auth0123").Return(&user, nil).Once()

	result, err := suite.service.GetAll(1)

	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), companiesDTO, result)
}

func (suite *CompanyServiceUnitTestsSuite) TestCompanyService_Approve_CompanyRequestDeclined() {
	request := dto.ApproveCompanyDTO{
		ID:      1,
		Approve: false,
	}

	suite.companyRepositoryMock.On("Approve", &request).Return(nil).Once()

	err := suite.service.Approve(&request)

	assert.Equal(suite.T(), nil, err)
}

func (suite *CompanyServiceUnitTestsSuite) TestCompanyService_Update_CompanyUpdated() {
	company := dto.CompanyUpdateDTO{
		ID:          1,
		Name:        "new name",
		Contact:     "new contact",
		Description: "new description",
	}

	companyEntity := model.Company{
		ID:           1,
		Name:         "name",
		Contact:      "contact",
		Description:  "description",
		OwnerAuth0ID: "1",
	}

	suite.companyRepositoryMock.On("GetByID", company.ID).Return(&companyEntity, nil).Once()

	toUpdate := mapper.CompanyUpdateDTOToCompany(&company)
	toUpdate.OwnerAuth0ID = companyEntity.OwnerAuth0ID

	forReturn := dto.CompanyResponseDTO{
		ID:          1,
		Name:        "new name",
		Contact:     "new contact",
		Description: "new description",
	}
	suite.companyRepositoryMock.On("Update", toUpdate).Return(&forReturn, nil).Once()

	updatedCompany, err := suite.service.Update(&company, "1")

	assert.Equal(suite.T(), company.Name, updatedCompany.Name)
	assert.Equal(suite.T(), company.Contact, updatedCompany.Contact)
	assert.Equal(suite.T(), company.Description, updatedCompany.Description)
	assert.Equal(suite.T(), nil, err)
}
