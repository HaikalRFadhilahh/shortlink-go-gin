package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/HaikalRFadhilahh/shortlink-go-gin/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tkn := ctx.GetHeader("Authorization")
		if tkn == "" {
			ctx.JSON(http.StatusForbidden, helper.ErrorReponse{
				StatusCode: http.StatusForbidden,
				Status:     "error",
				Message:    "Token Invalid!",
			})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tkn, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid encryption")
			}
			return []byte(helper.GetEnv("JWT_SECRET", "")), nil
		})
		if err != nil {
			ctx.JSON(http.StatusForbidden, helper.ErrorReponse{
				StatusCode: http.StatusForbidden,
				Status:     "error",
				Message:    "Token Invalid!",
			})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.JSON(http.StatusForbidden, helper.ErrorReponse{
					StatusCode: http.StatusForbidden,
					Status:     "error",
					Message:    "Token Invalid!",
				})
				ctx.Abort()
				return
			}
			ctx.Set("userId", claims["userId"])
		} else {
			ctx.JSON(http.StatusForbidden, helper.ErrorReponse{
				StatusCode: http.StatusForbidden,
				Status:     "error",
				Message:    "Token Invalid!",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
