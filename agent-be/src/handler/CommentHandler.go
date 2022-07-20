package handler

import (
	"agent-app/src/dto"
	"agent-app/src/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	Service *service.CommentService
}

func (handler *CommentHandler) AddComment(ctx *gin.Context) {
	var commentDTO dto.CommentDTO
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	commentId, err := handler.Service.AddComment(&commentDTO)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, commentId)
}

func (handler *CommentHandler) DeleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = handler.Service.DeleteComment(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}

func (handler *CommentHandler) UpdateComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var commentDTO dto.CommentDTO
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	commentResultDTO, err := handler.Service.UpdateComment(id, &commentDTO)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, commentResultDTO)
}

func (handler *CommentHandler) GetCommentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	comment, err := handler.Service.GetCommentById(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (handler *CommentHandler) GetCommentByOwnerID(ctx *gin.Context) {
	ownerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	comments, err := handler.Service.GetCommentByOwnerID(ownerId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (handler *CommentHandler) GetCommentByCompanyID(ctx *gin.Context) {
	companyId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	comments, err := handler.Service.GetCommentByCompanyID(companyId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
