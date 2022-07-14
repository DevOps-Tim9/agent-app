package service

import (
	"agent-app/src/dto"
	"agent-app/src/mapper"
	"agent-app/src/repository"
)

type ICommentService interface {
	AddComment(comment *dto.CommentDTO) (int, error)
	DeleteComment(int) error
	UpdateComment(int, *dto.CommentDTO) (int, error)
	GetCommentByOwnerID(int) (*[]dto.CommentDTO, error)
	GetCommentByCompanyID(int) (*[]dto.CommentDTO, error)
	GetCommentById(int) (*dto.CommentDTO, error)
}

type CommentService struct {
	CommentRepo repository.ICommentRepository
}

func NewCommentService(commentRepo repository.ICommentRepository) ICommentService {
	return &CommentService{
		commentRepo,
	}
}

func (service *CommentService) AddComment(companyDTO *dto.CommentDTO) (int, error) {
	commentID, err := service.CommentRepo.AddComment(mapper.CommentDTOToComment(companyDTO))
	if err != nil {
		return -1, err
	}
	return commentID, nil
}

func (service *CommentService) DeleteComment(id int) error {
	err := service.CommentRepo.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *CommentService) UpdateComment(id int, comment *dto.CommentDTO) (int, error) {
	result, err := service.CommentRepo.UpdateComment(id, mapper.CommentDTOToComment(comment))
	if err != nil {
		return -1, err
	}
	return result, nil
}

func (service *CommentService) GetCommentById(i int) (*dto.CommentDTO, error) {
	result, err := service.CommentRepo.GetCommentById(i)
	if err != nil {
		return nil, err
	}

	return mapper.CommentToCommentDTO(result), nil
}

func (service *CommentService) GetCommentByOwnerID(ownerId int) (*[]dto.CommentDTO, error) {
	result, err := service.CommentRepo.GetCommentByOwnerID(ownerId)
	if err != nil {
		return nil, err
	}

	return mapper.ListCommentTOListDTOs(*result), nil
}

func (service *CommentService) GetCommentByCompanyID(companyId int) (*[]dto.CommentDTO, error) {
	result, err := service.CommentRepo.GetCommentByCompanyID(companyId)
	if err != nil {
		return nil, err
	}

	return mapper.ListCommentTOListDTOs(*result), nil
}
