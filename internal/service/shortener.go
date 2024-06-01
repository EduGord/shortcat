package service

import (
	"errors"
	"net/url"
	"short.cat/internal/repository"
	"short.cat/pkg/shortener"
)

type ShortURLServiceImpl struct {
	repo repository.ShortURLRepository
}

func NewShortURLService(repo repository.ShortURLRepository) *ShortURLServiceImpl {
	return &ShortURLServiceImpl{repo: repo}
}

func (s *ShortURLServiceImpl) ShortenURL(originalURL string, keyword string) (string, error) {
	if err := s.validateOriginalURL(originalURL); err != nil {
		return "", err
	}

	if keyword == "" {
		var err error
		keyword, err = s.generateUniqueKeyword()
		if err != nil {
			return "", err
		}
	} else {
		if _, err := s.repo.Find(keyword); err == nil {
			return "", errors.New("conflict, keyword already exists")
		}
	}

	if err := s.repo.Save(keyword, originalURL); err != nil {
		return "", err
	}

	return keyword, nil
}

func (s *ShortURLServiceImpl) GetOriginalURL(shortURL string) (string, error) {
	return s.repo.Find(shortURL)
}

func (s *ShortURLServiceImpl) validateOriginalURL(originalURL string) error {
	if originalURL == "" {
		return errors.New("original URL is required")
	}
	if _, err := url.ParseRequestURI(originalURL); err != nil {
		return errors.New("invalid original URL")
	}
	return nil
}

func (s *ShortURLServiceImpl) generateUniqueKeyword() (string, error) {
	const maxRetries = 5

	for i := 0; i < maxRetries; i++ {
		keyword := shortener.GenerateShortURL()
		if _, err := s.repo.Find(keyword); err != nil {
			return keyword, nil
		}
	}

	return "", errors.New("failed to generate a unique keyword, please try again")
}
