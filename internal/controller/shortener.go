package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short.cat/internal/model"
	"short.cat/internal/service"
	"strings"
)

type ShortURLControllerImpl struct {
	service *service.ShortURLServiceImpl
}

func NewShortURLController(service *service.ShortURLServiceImpl) *ShortURLControllerImpl {
	return &ShortURLControllerImpl{service: service}
}

func (ctrl *ShortURLControllerImpl) ShortenURL(c *gin.Context) {
	var request model.ShortenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL, err := ctrl.service.ShortenURL(request.URL, request.Keyword)
	if err != nil {
		if strings.Contains(err.Error(), "conflict") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := model.ShortenResponse{ShortURL: shortURL}
	c.JSON(http.StatusCreated, response)
}

func (ctrl *ShortURLControllerImpl) Redirect(c *gin.Context) {
	shortURL := c.Param("shortURL")

	originalURL, err := ctrl.service.GetOriginalURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}
