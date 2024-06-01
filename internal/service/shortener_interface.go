package service

type ShortURLService interface {
	ShortenURL(originalURL string, keyword string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
}
