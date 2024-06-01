package controller

import (
	"github.com/gin-gonic/gin"
)

type ShortURLController interface {
	ShortenURL(c *gin.Context)
	Redirect(c *gin.Context)
}
