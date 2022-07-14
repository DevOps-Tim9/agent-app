package handler

import (
	"agent-app/src/dto"
	"agent-app/src/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var userToRegister dto.RegistrationRequestDTO
	if err := ctx.ShouldBindJSON(&userToRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	userID, err := handler.Service.Register(&userToRegister)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, userID)
}

func (handler *UserHandler) GetByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	user, err := handler.Service.GetByEmail(email)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}
