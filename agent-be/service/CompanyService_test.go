package service

import (
	"agent-app/auth0"
	"agent-app/dto"
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
}

func TestCompanyServiceUnitTestsSuite(t *testing.T) {
	suite.Run(t, new(CompanyServiceUnitTestsSuite))
}

func (suite *CompanyServiceUnitTestsSuite) SetupSuite() {
	suite.companyRepositoryMock = new(repository.CompanyRepositoryMock)
	suite.auth0ClientMock = new(auth0.Auth0ClientMock)
	suite.service = NewCompanyService(suite.companyRepositoryMock, suite.auth0ClientMock)
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
	}
	var companies []*model.Company
	companies = append(companies, &company)

	var companiesDTO []*dto.CompanyResponseDTO
	companiesDTO = append(companiesDTO, &companyDTO)

	suite.companyRepositoryMock.On("GetAll", 1).Return(companies, nil).Once()

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
