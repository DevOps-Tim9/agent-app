package repository

import (
	"agent-app/src/dto"
	"agent-app/src/mapper"
	"agent-app/src/model"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ICompanyRepository interface {
	AddCompany(*model.Company) (int, error)
	Approve(*dto.ApproveCompanyDTO) error
	GetAll(int) ([]*model.Company, error)
	GetByID(int) (*model.Company, error)
	Update(*model.Company) (*dto.CompanyResponseDTO, error)
}

func NewCompanyRepository(database *gorm.DB) ICompanyRepository {
	return &CompanyRepository{
		database,
	}
}

type CompanyRepository struct {
	Database *gorm.DB
}

func (repo *CompanyRepository) AddCompany(company *model.Company) (int, error) {
	result := repo.Database.Create(company)

	if result.Error != nil {
		return -1, result.Error
	}

	return company.ID, nil
}

func (repo *CompanyRepository) Approve(approveCompanyDTO *dto.ApproveCompanyDTO) error {
	companyEntity := model.Company{
		ID: approveCompanyDTO.ID,
	}
	if err := repo.Database.Where("ID = ?", approveCompanyDTO.ID).First(&companyEntity).Error; err != nil {
		return errors.New(fmt.Sprintf("Company with ID %d not found", approveCompanyDTO.ID))
	}

	if approveCompanyDTO.Approve {
		companyEntity.Approved = true
		result := repo.Database.Save(&companyEntity)
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := repo.Database.Delete(&model.Company{}, approveCompanyDTO.ID)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (repo *CompanyRepository) GetAll(approved int) ([]*model.Company, error) {
	var companies = []*model.Company{}
	if result := repo.Database.Find(&companies, "Approved = ?", approved); result.Error != nil {
		return nil, errors.New("Error happened during retrieving entities from database")
	}

	return companies, nil
}

func (repo *CompanyRepository) GetByID(ID int) (*model.Company, error) {
	var company = model.Company{}
	if result := repo.Database.Find(&company, "ID = ?", ID); result.Error != nil {
		return nil, result.Error
	}

	return &company, nil
}

func (repo *CompanyRepository) Update(company *model.Company) (*dto.CompanyResponseDTO, error) {
	result := repo.Database.Save(&company)

	if result.Error != nil {
		return nil, result.Error
	}

	return mapper.CompanyToCompanyResponseDTO(company), nil
}
