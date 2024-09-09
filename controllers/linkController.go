package controllers

import (
	"net/http"

	"github.com/HaikalRFadhilahh/shortlink-go-gin/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LinkController struct {
	DB *gorm.DB
}

func (l *LinkController) CreateLink(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, helper.LinkResponse{
		StatusCode: http.StatusCreated,
		Status:     "success",
		Message:    "Link Created",
	})
}
