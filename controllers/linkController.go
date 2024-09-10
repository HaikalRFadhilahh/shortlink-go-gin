package controllers

import (
	"net/http"

	"github.com/HaikalRFadhilahh/shortlink-go-gin/helper"
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
