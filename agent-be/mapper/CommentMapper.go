package mapper

import (
	"agent-app/dto"
	"agent-app/model"
)

func CommentDTOToComment(dto *dto.CommentDTO) *model.Comment {
	comment := model.Comment{
		ID:               dto.ID,
		CreationDate:     dto.CreationDate,
		UserOwnerID:      dto.UserOwnerID,
		CompanyID:        dto.CompanyID,
		Position:         dto.Position,
		Salary:           dto.Salary,
		Rating:           dto.Rating,
		InterviewProcess: dto.InterviewProcess,
		Description:      dto.Description,
	}
	return &comment
}

func CommentToCommentDTO(comment *model.Comment) *dto.CommentDTO {
	commentDTO := dto.CommentDTO{
		ID:               comment.ID,
		CreationDate:     comment.CreationDate,
		UserOwnerID:      comment.UserOwnerID,
		CompanyID:        comment.CompanyID,
		Position:         comment.Position,
		Salary:           comment.Salary,
		Rating:           comment.Rating,
		InterviewProcess: comment.InterviewProcess,
		Description:      comment.Description,
	}
	return &commentDTO
}

func ListCommentTOListDTOs(comments []model.Comment) *[]dto.CommentDTO {
	var dtoList []dto.CommentDTO
	for _, element := range comments {
		dtoList = append(dtoList, *CommentToCommentDTO(&element))
	}
	return &dtoList
}
