package repository

import (
	"agent-app/dto"
	"agent-app/model"

	"github.com/stretchr/testify/mock"
)

type CompanyRepositoryMock struct {
	mock.Mock
}

func (c *CompanyRepositoryMock) AddCompany(company *model.Company) (int, error) {
	args := c.Called(company)

	return args.Int(0), args.Error(1)
}

func (c *CompanyRepositoryMock) Approve(approveCompanyDTO *dto.ApproveCompanyDTO) error {
	args := c.Called(approveCompanyDTO)

	return args.Error(0)
}

func (c *CompanyRepositoryMock) GetAll(approved int) ([]*model.Company, error) {
	args := c.Called(approved)

	return args.Get(0).([]*model.Company), args.Error(1)
}
