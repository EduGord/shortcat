package service

import (
	"github.com/go-playground/assert/v2"
	"short.cat/internal/repository"
	"short.cat/internal/service"
	"testing"
)

func TestShortenURL(t *testing.T) {
	r := repository.NewShortURLRepository()
	s := service.NewShortURLService(r)

	originalURL := "https://example.com"
	shortURL, err := s.ShortenURL(originalURL, "example")

	assert.Equal(t, err, nil)
	assert.Equal(t, shortURL, "example")

	shortURL, err = s.ShortenURL(originalURL, "")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, shortURL, "")
}

func TestGetOriginalURL(t *testing.T) {
	r := repository.NewShortURLRepository()
	s := service.NewShortURLService(r)

	originalURL := "https://example.com"
	shortURL, _ := s.ShortenURL(originalURL, "example")

	foundURL, err := s.GetOriginalURL(shortURL)
	assert.Equal(t, err, nil)
	assert.Equal(t, foundURL, originalURL)
}
