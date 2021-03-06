package user_route

import (
	"github.com/gin-gonic/gin"
	"github.com/sideneckx/template-go-server/api_model"
	"github.com/sideneckx/template-go-server/auth/auth_middleware"
	"github.com/sideneckx/template-go-server/model"
)

var userRoute *gin.RouterGroup = nil

func NewRoute(engine gin.Engine) {
	userRoute = engine.Group("/api/user")

	userRoute.GET("", get)
	userRoute.POST("/", create)
	userRoute.GET("/require-token", auth_middleware.AuthorizeMiddleWare(), requiredSignedIn)
	userRoute.GET("/:id", getId)
}

func get(c *gin.Context) {

	name := c.Query("name")
	if name == "" {
		c.JSON(400, api_model.NewErrorAPIRes("name is invalid or missing", api_model.MissingParam))
		return
	}
	mockUser := model.User{Id: 1, Name: name}
	json := api_model.NewSuccessAPIRes(&mockUser)
	c.JSON(200, json.ToMap())
}

func create(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(400, api_model.NewErrorAPIRes("name is required", api_model.MissingParam))
		return
	}
	mockUser := model.User{Id: 2, Name: name}
	json := api_model.NewSuccessAPIRes(&mockUser)
	c.JSON(200, json.ToMap())
}

func getId(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, api_model.NewErrorAPIRes("id is required", api_model.MissingParam))
		return
	}
	if id == "1" {
		mockUser := model.User{Id: 2, Name: "Laura"}
		json := api_model.NewSuccessAPIRes(&mockUser)
		c.JSON(200, json.ToMap())
		return
	}
	c.JSON(400, api_model.NewErrorAPIRes("id not found", api_model.RequestFail))
}

func requiredSignedIn(c *gin.Context) {
	msg := "Passed the middleware"
	c.JSON(200, api_model.NewSuccessAPIRes(&msg))
}
