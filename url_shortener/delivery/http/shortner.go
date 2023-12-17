package http

import (
	"fmt"
	"net/http"
	"url-shortener/domain"

	"github.com/gin-gonic/gin"
)

type URLShortenerHttpHandler struct {
	service domain.URLShortenerService
}

// returns a new http handler for url shortener api
func NewURLShortenerHttpHandler(srvc domain.URLShortenerService) *URLShortenerHttpHandler {
	return &URLShortenerHttpHandler{service: srvc}
}

// ShortenURL godoc
// @Summary Shorten URL API
// @Description Returns a shorten URL for input URL
// @Tags URLShortener
// @Produce json
// @Param request body ShortenURLRequest true "json with actual url"
// @Success 200 {object} ShortenURLSuccessResponse
// @Failure 400 {object} ShortenURLErrorResponse
// @Failure 500 {object} ShortenURLErrorResponse
// @Router /url/shorten [post]
func (h *URLShortenerHttpHandler) ShortenURL(c *gin.Context) {
	req := ShortenURLRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ShortenURLErrorResponse{Message: err.Error()})
		return
	}

	shortURL, err := h.service.ShortenURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ShortenURLErrorResponse{Error: err.Error(), Message: "FAILED"})
		return
	}

	scheme := "http"
	if c.Request.URL.Scheme != "" {
		scheme = c.Request.URL.Scheme
	}
	url := fmt.Sprintf(domain.ShortURLFormat, scheme, c.Request.Host, shortURL)

	c.JSON(http.StatusOK, ShortenURLSuccessResponse{Message: "SUCCESS", ShortURL: url})
}

// RedirectToFullURL godoc
// @Summary Redirector API
// @Description Redirects the shortened URL to actual URL location
// @Tags URLShortener
// @Param key path string true "short url code"
// @Success 301
// @Failure 500 {object} ShortenURLErrorResponse
// @Router /{key} [get]
func (h *URLShortenerHttpHandler) RedirectToFullURL(c *gin.Context) {
	hash := c.Param("key")

	fullURL, err := h.service.GetOriginalURL(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ShortenURLErrorResponse{Error: err.Error(), Message: "FAILED"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, fullURL)
}
