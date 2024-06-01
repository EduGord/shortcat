package repository

import "errors"

type inMemoryRepository struct {
	urls map[string]string
}

func NewShortURLRepository() ShortURLRepository {
	return &inMemoryRepository{urls: make(map[string]string)}
}

func (r *inMemoryRepository) Save(shortURL string, originalURL string) error {
	if _, exists := r.urls[shortURL]; exists {
		return errors.New("short URL already exists")
	}
	r.urls[shortURL] = originalURL
	return nil
}

func (r *inMemoryRepository) Find(shortURL string) (string, error) {
	originalURL, exists := r.urls[shortURL]
	if !exists {
		return "", errors.New("short URL not found")
	}
	return originalURL, nil
}
