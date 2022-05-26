package api_route

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sideneckx/template-go-server/api_model"
)

func ThrowError(c *gin.Context, err error) {
	c.JSON(400, api_model.NewErrorAPIRes(err.Error(), api_model.RequestFail))
}

func SendJsonData(c *gin.Context, data any) {
	c.JSON(200, api_model.NewSuccessAPIRes(&data).ToMap())
}

func MakeSureParam(c *gin.Context, method string, params []string) error {
	err := errors.New("Exist error")
	for _, param := range params {
		isSatisfy := false
		if method == http.MethodGet {
			isSatisfy = c.Query(param) != ""
		}
		if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
			isSatisfy = c.PostForm(param) != ""
		}

		if !isSatisfy {
			c.JSON(400, api_model.NewErrorAPIRes(param+" is required", api_model.MissingParam))
			return err
		}
	}
	return nil
}

func MakeSureParamOnUrl(c *gin.Context, params []string) error {
	err := errors.New("Exist error")
	for _, param := range params {
		isSatisfy := c.Param(param) != ""
		if !isSatisfy {
			c.JSON(400, api_model.NewErrorAPIRes(param+" is required", api_model.MissingParam))
			return err
		}
	}
	return nil
}
