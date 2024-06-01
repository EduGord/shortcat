package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"short.cat/internal/controller"
	"short.cat/internal/repository"
	"short.cat/internal/service"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setupRouter() *gin.Engine {
	r := repository.NewShortURLRepository()
	s := service.NewShortURLService(r)
	c := controller.NewShortURLController(s)

	router := gin.Default()
	router.POST("/", c.ShortenURL)
	router.GET("/:shortURL", c.Redirect)

	return router
}

func TestShortenURLNoKeywordSuccess(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"url":"https://example.com"}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestShortenURLNoBodyFailure(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestShortenURLWithKeyword(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	body := `{"url":"https://example.com", "keyword":"example"}`
	req.Body = io.NopCloser(strings.NewReader(body))

	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, w.Body.String(), `{"short_url":"example"}`)
}

func TestRedirect(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", http.NoBody)
	req.Header.Set("Content-Type", "application/json")

	body := `{"url":"https://example.com", "keyword":"example"}`
	req.Body = io.NopCloser(strings.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusCreated)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/example", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusFound)
	assert.Equal(t, w.Header().Get("Location"), "https://example.com")
}

func TestNotFound(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/notfound", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)
	assert.Equal(t, w.Body.String(), `{"error":"URL not found"}`)
}

func TestConflictKeyword(t *testing.T) {
	router := setupRouter()

	for i := 0; i < 2; i++ {
		w := httptest.NewRecorder()
		body := `{"url":"https://example.com", "keyword":"example"}`
		req, _ := http.NewRequest("POST", "/", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		if i == 0 {
			assert.Equal(t, http.StatusCreated, w.Code)
		} else {
			assert.Equal(t, http.StatusConflict, w.Code)
			assert.Equal(t, `{"error":"conflict, keyword already exists"}`, w.Body.String())
		}
	}
}
