package api_route

import (
	"github.com/gin-gonic/gin"
	"github.com/sideneckx/template-go-server/api_model"
)

func ThrowError(c *gin.Context, err error) {
	c.JSON(400, api_model.NewErrorAPIRes(err.Error(), api_model.RequestFail))
}

func SendJsonData(c *gin.Context, data any) {
	c.JSON(200, api_model.NewSuccessAPIRes(&data).ToMap())
}
