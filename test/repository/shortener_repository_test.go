package repository

import (
	"github.com/go-playground/assert/v2"
	"short.cat/internal/repository"
	"testing"
)

func newTestRepository() repository.ShortURLRepository {
	return repository.NewShortURLRepository()
}

func TestSaveNewShortURL(t *testing.T) {
	repo := newTestRepository()

	err := repo.Save("short1", "http://original-url.com")
	assert.Equal(t, err, nil)
}

func TestSaveExistingShortURL(t *testing.T) {
	repo := newTestRepository()

	err := repo.Save("short1", "http://original-url.com")
	assert.Equal(t, err, nil)

	err = repo.Save("short1", "http://another-url.com")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "short URL already exists")
}

func TestFindNonExistentShortURL(t *testing.T) {
	repo := newTestRepository()

	_, err := repo.Find("nonexistent")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "short URL not found")
}

func TestFindExistingShortURL(t *testing.T) {
	repo := newTestRepository()

	err := repo.Save("short1", "http://original-url.com")
	assert.Equal(t, err, nil)

	originalURL, err := repo.Find("short1")
	assert.Equal(t, err, nil)
	assert.Equal(t, originalURL, "http://original-url.com")
}

func TestSaveAndFindIntegration(t *testing.T) {
	repo := newTestRepository()

	err := repo.Save("short2", "http://example.com")
	assert.Equal(t, err, nil)

	originalURL, err := repo.Find("short2")
	assert.Equal(t, err, nil)
	assert.Equal(t, originalURL, "http://example.com")
}
