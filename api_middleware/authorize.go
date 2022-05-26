package api_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sideneckx/template-go-server/api_model"
)

func AuthorizeMiddleWare(exceptPath ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, path := range exceptPath {
			if c.FullPath() == path {
				c.Next()
				return
			}
		}
		code := c.GetHeader("Secret-code")
		if code != "uwu" {
			sendNeedAuthorizeRes(c)
			return
		}
		c.Next()
	}
}

func sendNeedAuthorizeRes(c *gin.Context) {
	c.AbortWithStatusJSON(400, api_model.NewErrorAPIRes("Secret code is invalid", api_model.NeedAuthorize).ToMap())
}
