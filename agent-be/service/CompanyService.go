package service

import (
	"agent-app/auth0"
	"agent-app/dto"
	"agent-app/mapper"
	"agent-app/repository"
	"fmt"
)

type CompanyService struct {
	CompanyRepo repository.ICompanyRepository
	Auth0Client auth0.Auth0Client
}

type ICompanyService interface {
	Register(*dto.CompanyRequestDTO, string) (int, error)
	Approve(*dto.ApproveCompanyDTO) error
	GetAll(approved int) ([]*dto.CompanyResponseDTO, error)
}

func NewCompanyService(companyRepository repository.ICompanyRepository, auth0Client auth0.Auth0Client) ICompanyService {
	return &CompanyService{
		companyRepository,
		auth0Client,
	}
}

func (service *CompanyService) Register(companyToRegister *dto.CompanyRequestDTO, ownerAuth0ID string) (int, error) {
	company := mapper.CompanyRequestDTOToCompany(companyToRegister)
	company.OwnerAuth0ID = ownerAuth0ID

	err := company.Validate()
	if err != nil {
		return -1, err
	}

	addedCompanyID, err := service.CompanyRepo.AddCompany(company)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return addedCompanyID, nil
}

func (service *CompanyService) Approve(approveCompanyDTO *dto.ApproveCompanyDTO) error {
	return service.CompanyRepo.Approve(approveCompanyDTO)
}

func (service *CompanyService) GetAll(approved int) ([]*dto.CompanyResponseDTO, error) {
	companies, err := service.CompanyRepo.GetAll(approved)

	if err != nil {
		return nil, err
	}

	res := make([]*dto.CompanyResponseDTO, len(companies))
	for i := 0; i < len(companies); i++ {
		res[i] = mapper.CompanyToCompanyResponseDTO(companies[i])
	}

	return res, nil
}
