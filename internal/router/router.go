package router

import (
	"github.com/gin-gonic/gin"
	"short.cat/internal/controller"
	"short.cat/internal/repository"
	"short.cat/internal/service"
)

func SetupRouter() *gin.Engine {
	r := repository.NewShortURLRepository()
	s := service.NewShortURLService(r)
	c := controller.NewShortURLController(s)

	router := gin.Default()
	setupShortUrlRoutes(router, c)

	return router
}

func setupShortUrlRoutes(router *gin.Engine, c *controller.ShortURLControllerImpl) {
	router.POST("/", c.ShortenURL)
	router.GET("/:shortURL", c.Redirect)
}
