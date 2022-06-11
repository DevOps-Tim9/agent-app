package handler

import (
	"agent-app/dto"
	"agent-app/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	Service *service.CompanyService
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	token, _ := jwt.Parse(strings.Split(tokenStr, " ")[1], nil)

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, true
	} else {
		fmt.Println("Invalid JWT Token")
		return nil, false
	}
}

func (handler *CompanyHandler) Register(ctx *gin.Context) {
	claims, isValid := extractClaims(ctx.Request.Header.Get("Authorization"))
	if !isValid {
		fmt.Println("Not valid")
		ctx.JSON(http.StatusBadRequest, "Invalid token")
		return
	}

	var companyToRegister dto.CompanyRequestDTO
	if err := ctx.ShouldBindJSON(&companyToRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	companyID, err := handler.Service.Register(&companyToRegister, fmt.Sprint(claims["sub"]))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, companyID)
}

func (handler *CompanyHandler) Approve(ctx *gin.Context) {
	var approveCompanyDTO dto.ApproveCompanyDTO
	if err := ctx.ShouldBindJSON(&approveCompanyDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err := handler.Service.Approve(&approveCompanyDTO)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, true)
}

func (handler *CompanyHandler) GetAllCompanies(ctx *gin.Context) {
	companies, err := handler.Service.GetAll(1)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

func (handler *CompanyHandler) GetAllCompanyRequests(ctx *gin.Context) {
	companies, err := handler.Service.GetAll(0)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

func (handler *CompanyHandler) UpdateCompany(ctx *gin.Context) {
	var companyToUpdate dto.CompanyUpdateDTO
	if err := ctx.ShouldBindJSON(&companyToUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	claims, _ := extractClaims(ctx.Request.Header.Get("Authorization"))

	companyDTO, err := handler.Service.Update(&companyToUpdate, fmt.Sprint(claims["sub"]))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, companyDTO)
}
