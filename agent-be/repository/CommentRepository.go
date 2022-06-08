package repository

import (
	"agent-app/model"
	"github.com/jinzhu/gorm"
	"time"
)

type ICommentRepository interface {
	AddComment(comment *model.Comment) (int, error)
	DeleteComment(int) error
	UpdateComment(int, *model.Comment) (int, error)
	GetCommentByOwnerID(int) (*[]model.Comment, error)
	GetCommentByCompanyID(int) (*[]model.Comment, error)
	GetCommentById(int) (*model.Comment, error)
}

func NewCommentRepository(database *gorm.DB) ICommentRepository {
	return &CommentRepository{
		database,
	}
}

type CommentRepository struct {
	Database *gorm.DB
}

func (repo *CommentRepository) AddComment(company *model.Comment) (int, error) {
	company.CreationDate = time.Now()
	result := repo.Database.Create(company)

	if result.Error != nil {
		return -1, result.Error
	}

	return company.ID, nil
}

func (repo *CommentRepository) DeleteComment(id int) error {
	result := repo.Database.Delete(&model.Comment{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *CommentRepository) UpdateComment(id int, comment *model.Comment) (int, error) {
	comment.ID = id
	updatedId := repo.Database.Save(&comment)

	if updatedId.Error != nil {
		return -1, updatedId.Error
	}

	return id, nil
}

func (repo *CommentRepository) GetCommentById(id int) (*model.Comment, error) {
	var comment model.Comment
	result := repo.Database.Find(&comment, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (repo *CommentRepository) GetCommentByOwnerID(ownerId int) (*[]model.Comment, error) {
	var comments []model.Comment
	result := repo.Database.Where("user_owner_id = ?", ownerId).Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}

	return &comments, nil
}

func (repo *CommentRepository) GetCommentByCompanyID(companyId int) (*[]model.Comment, error) {
	var comments []model.Comment
	result := repo.Database.Where("company_id = ?", companyId).Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}

	return &comments, nil
}
