package controllers

import (
	"net/http"

	"github.com/HaikalRFadhilahh/shortlink-go-gin/helper"
	"github.com/HaikalRFadhilahh/shortlink-go-gin/models"
	"github.com/HaikalRFadhilahh/shortlink-go-gin/requestHandler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LinkController struct {
	DB *gorm.DB
}

func (l *LinkController) CreateLink(ctx *gin.Context) {
	requestLinkHandler := requestHandler.CreateLinkHeader{}
	userId, _ := ctx.Get("userId")
	requestLinkHandler.IsActive = true

	err := ctx.ShouldBindBodyWithJSON(&requestLinkHandler)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.LinkResponse{
			StatusCode: http.StatusBadRequest,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}

	var checkLink int64
	l.DB.Table("links").Where("alias=?", requestLinkHandler.Alias).Count(&checkLink)
	if checkLink > 0 {
		ctx.JSON(http.StatusConflict, helper.LinkResponse{
			StatusCode: http.StatusConflict,
			Status:     "error",
			Message:    "Shortlink Exist",
		})
		return
	}

	requestLinkHandler.UserId, _ = userId.(*int)

	err = l.DB.Table("links").Create(&requestLinkHandler).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.LinkResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, helper.LinkResponse{
		StatusCode: http.StatusCreated,
		Status:     "success",
		Message:    "Link Created",
		Data:       requestLinkHandler,
	})
}

func (u *LinkController) DeleteLink(ctx *gin.Context) {
	idLink := ctx.Param("idLink")
	link := models.Link{}

	checkLink := u.DB.Table("links").Where("id=?", idLink).Find(&link).RowsAffected > 0
	if !checkLink {
		ctx.JSON(http.StatusNotFound, helper.LinkResponse{
			StatusCode: http.StatusNotFound,
			Status:     "error",
			Message:    "Data Links Not Found!",
		})
		return
	}

	err := u.DB.Table("links").Delete(&link).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.LinkResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, helper.LinkResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Data link success deleted",
	})
}

func (u *LinkController) UpdateLink(ctx *gin.Context) {

}

func (u *LinkController) GetAllLink(ctx *gin.Context) {
	links := []models.Link{}

	err := u.DB.Table("links").Find(&links).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.LinkResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, helper.LinkResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "Data All Links",
		Data:       links,
	})
}

func (u *LinkController) GetRedirectLink(ctx *gin.Context) {
	alias := ctx.Param("alias")
	link := models.Link{}

	checkLink := u.DB.Table("links").Where("alias=?", alias).First(&link).RowsAffected == 1
	if !checkLink {
		ctx.JSON(http.StatusNotFound, helper.LinkResponse{
			StatusCode: http.StatusNotFound,
			Status:     "error",
			Message:    "Link Not Found!",
		})
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, link.Link)
}
