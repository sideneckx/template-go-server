package auth_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sideneckx/template-go-server/api_model"
	"github.com/sideneckx/template-go-server/auth/server_jwt"
)

func AuthorizeMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if _, err := server_jwt.DecodeTokenUserId(token); err != nil {
			sendNeedAuthorizeRes(c)
			return
		}
		c.Next()
	}
}

func sendNeedAuthorizeRes(c *gin.Context) {
	c.AbortWithStatusJSON(400, api_model.NewErrorAPIRes("Token in header is invalid or missing", api_model.NeedAuthorize).ToMap())
}
