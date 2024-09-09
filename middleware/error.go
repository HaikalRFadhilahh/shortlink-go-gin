package middleware

import (
	"net/http"

	"github.com/HaikalRFadhilahh/shortlink-go-gin/helper"
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(http.StatusInternalServerError, helper.ErrorReponse{
					StatusCode: 500,
					Status:     "error",
					Message:    "Internal Server Error!",
				})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
