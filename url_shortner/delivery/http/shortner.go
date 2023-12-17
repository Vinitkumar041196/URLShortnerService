package http

import (
	"fmt"
	"net/http"
	"url-shortner/domain"

	"github.com/gin-gonic/gin"
)

type URLShortnerHttpHandler struct {
	service domain.URLShortnerService
}

//returns a new http handler for url shortner api
func NewURLShortnerHttpHandler(srvc domain.URLShortnerService) *URLShortnerHttpHandler {
	return &URLShortnerHttpHandler{service: srvc}
}

// ShortenURL godoc
// @Summary Shorten URL API
// @Description Returns a shorten URL for input URL
// @Tags URLShortner
// @Produce json
// @Param request body ShortenURLRequest true "json with actual url"
// @Success 200 {object} ShortenURLSuccessResponse
// @Failure 400 {object} ShortenURLErrorResponse
// @Failure 500 {object} ShortenURLErrorResponse
// @Router /url/shorten [post]
func (h *URLShortnerHttpHandler) ShortenURL(c *gin.Context) {
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

	url := fmt.Sprintf(domain.ShortURLFormat, c.Request.Host, shortURL)

	c.JSON(http.StatusOK, ShortenURLSuccessResponse{Message: "SUCCESS", ShortURL: url})
}

// RedirectToFullURL godoc
// @Summary Redirector API
// @Description Redirects the shortened URL to actual URL location
// @Tags URLShortner
// @Param key path string true "short url code"
// @Success 301
// @Failure 500 {object} ShortenURLErrorResponse
// @Router /{key} [get]
func (h *URLShortnerHttpHandler) RedirectToFullURL(c *gin.Context) {
	hash := c.Param("key")

	fullURL, err := h.service.GetOriginalURL(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ShortenURLErrorResponse{Error: err.Error(), Message: "FAILED"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, fullURL)
}
