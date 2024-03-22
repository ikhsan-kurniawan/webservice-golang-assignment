package middlewares

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := strconv.Atoi(ctx.Param("userId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)

		if userID != int(userData["id"].(float64)) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				// "paramid": userID,
				// "userid": int(userData["id"].(float64)),
				"error":   "Unauthorized",
				"message": "kamu ga boleh akses data ini",
			})
			return
		}
		ctx.Next()
	}
}