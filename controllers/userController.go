package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/HaikalRFadhilahh/shortlink-go-gin/helper"
	"github.com/HaikalRFadhilahh/shortlink-go-gin/models"
	"github.com/HaikalRFadhilahh/shortlink-go-gin/requestHandler"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (u *UserController) Register(ctx *gin.Context) {
	registerHandler := requestHandler.RegisterHandler{}
	err := ctx.ShouldBindBodyWithJSON(&registerHandler)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.UserResponse{
			StatusCode: http.StatusBadRequest,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}

	// Username and Email Validation
	var checkCredentials int64
	u.DB.Table("users").Where("email=?", registerHandler.Email).Count(&checkCredentials)
	if checkCredentials > 0 {
		ctx.JSON(http.StatusConflict, helper.UserResponse{
			StatusCode: http.StatusConflict,
			Status:     "error",
			Message:    "Email Exist",
		})
		return
	}
	u.DB.Table("users").Where("username=?", registerHandler.Username).Count(&checkCredentials)
	if checkCredentials > 0 {
		ctx.JSON(http.StatusConflict, helper.UserResponse{
			StatusCode: http.StatusConflict,
			Status:     "error",
			Message:    "Username Exist",
		})
		return
	}

	res, err := bcrypt.GenerateFromPassword([]byte(registerHandler.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.UserResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}
	registerHandler.Password = string(res)

	// Register
	err = u.DB.Table("users").Create(&registerHandler).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.UserResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, helper.UserResponse{
		StatusCode: http.StatusCreated,
		Status:     "success",
		Message:    "register success",
		Data:       registerHandler,
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	loginHandler := requestHandler.LoginHandler{}
	user := models.User{}
	err := ctx.ShouldBindBodyWithJSON(&loginHandler)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.UserResponse{
			StatusCode: http.StatusBadRequest,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}

	check := u.DB.Table("users").Where("username=?", loginHandler.Username).First(&user).RowsAffected > 0
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(loginHandler.Password))

	if err != nil || !check {
		ctx.JSON(http.StatusUnauthorized, helper.UserResponse{
			StatusCode: http.StatusUnauthorized,
			Status:     "error",
			Message:    "username or password invalid!",
		})
		return
	}

	expTime, _ := strconv.Atoi((helper.GetEnv("JWT_EXPIRED_MINUTE", "1")))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
		"exp":    time.Now().Add(time.Minute * time.Duration(expTime)).Unix(),
		"iat":    time.Now().Unix(),
	})

	signedToken, err := token.SignedString([]byte(helper.GetEnv("JWT_SECRET", "")))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.UserResponse{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, helper.UserResponse{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    "User Authenticated!",
		Data: map[string]interface{}{
			"token": signedToken,
		},
	})
}
