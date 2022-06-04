package mapper

import (
	"agent-app/src/dto"
	"agent-app/src/model"
)

func CompanyRequestDTOToCompany(companyRequestDto *dto.CompanyRequestDTO) *model.Company {
	var company model.Company

	company.Name = companyRequestDto.Name
	company.Contact = companyRequestDto.Contact
	company.Description = companyRequestDto.Description

	return &company
}

func CompanyToCompanyResponseDTO(company *model.Company) *dto.CompanyResponseDTO {
	var companyDTO dto.CompanyResponseDTO

	companyDTO.ID = company.ID
	companyDTO.Name = company.Name
	companyDTO.Contact = company.Contact
	companyDTO.Description = company.Description

	return &companyDTO
}
