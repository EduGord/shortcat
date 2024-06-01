package model

type ShortenRequest struct {
	URL     string `json:"url" binding:"required"`
	Keyword string `json:"keyword"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
