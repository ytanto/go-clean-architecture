package infrastructure

import (
	gin "github.com/gin-gonic/gin"

	"github.com/ytanto/go-clean-architecture/pattern1/src/app/interfaces/controllers"
)

var Router *gin.Engin

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
