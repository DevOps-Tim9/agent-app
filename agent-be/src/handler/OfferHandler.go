package handler

import (
	"agent-app/src/dto"
	"agent-app/src/service"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OfferHandler struct {
	Service *service.OfferService
}

func (handler *OfferHandler) AddJobOffer(ctx *gin.Context) {
	var jobOfferDTO dto.JobOfferRequestDTO
	if err := ctx.ShouldBindJSON(&jobOfferDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	claims, isValid := extractClaims(ctx.Request.Header.Get("Authorization"))
	if !isValid {
		fmt.Println("Not valid")
		ctx.JSON(http.StatusBadRequest, "Invalid token")
		return
	}

	dto, err := handler.Service.Add(&jobOfferDTO, fmt.Sprint(claims["sub"]))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, dto)
}

func (handler *OfferHandler) GetAll(ctx *gin.Context) {
	offersDTO, err := handler.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *OfferHandler) GetJobOffersByCompany(ctx *gin.Context) {
	id, idErr := getId(ctx.Param("companyId"))
	if idErr != nil {
		ctx.JSON(http.StatusBadRequest, idErr.Error())
		return
	}

	offersDTO, err := handler.Service.GetCompanysOffers(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *OfferHandler) GetJobOfferById(ctx *gin.Context) {
	id, idErr := getId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(http.StatusBadRequest, idErr.Error())
		return
	}

	offersDTO, err := handler.Service.GetJobOfferById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *OfferHandler) DeleteJobOfferById(ctx *gin.Context) {
	id, idErr := getId(ctx.Param("id"))
	if idErr != nil {
		ctx.JSON(http.StatusBadRequest, idErr.Error())
		return
	}

	claims, isValid := extractClaims(ctx.Request.Header.Get("Authorization"))
	if !isValid {
		fmt.Println("Not valid")
		ctx.JSON(http.StatusBadRequest, "Invalid token")
		return
	}

	err := handler.Service.DeleteJobOffer(id, fmt.Sprint(claims["sub"]))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (handler *OfferHandler) Search(ctx *gin.Context) {
	param := ctx.Query("param")
	offersDTO, err := handler.Service.Search(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func getId(idParam string) (int, error) {
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return 0, errors.New("Company id should be a number")
	}
	return int(id), nil
}
