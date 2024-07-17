package middlewares

import (
	"AvaImageServer/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthByAvaImgServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		remotePassword := c.Request.Header.Get("Authorization")
		localPassword := lib.MD5WithSalt()
		if remotePassword == localPassword {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无授权"})
			c.Abort()
		}
	}
}
