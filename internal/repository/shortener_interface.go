package repository

type ShortURLRepository interface {
	Save(shortURL string, originalURL string) error
	Find(shortURL string) (string, error)
}
