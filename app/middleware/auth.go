package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/base"
	errEnum "github.com/gin-qt-business/app/errors"
	"github.com/gin-qt-business/app/utils/jwt"
)

const HEADER_TOKEN = "Authorization"

func AccessCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get(HEADER_TOKEN)
		if token != "" {
			jwt := &jwt.JWT{
				Token: token,
			}
			valid, err := jwt.ValidateToken()
			if err != nil {
				context.AbortWithStatusJSON(http.StatusUnauthorized, base.NewResponseError(err.Error()))
			}
			if !valid {
				context.AbortWithStatusJSON(http.StatusUnauthorized, base.NewResponseError(errEnum.TOKEN_INVALID_ERROR))
			}
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, base.NewResponseError(errEnum.TOKEN_INVALID_ERROR))
		}
		context.Next()
	}
}
