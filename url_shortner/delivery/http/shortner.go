package http

import (
	"fmt"
	"net/http"
	"url-shortner/domain"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	service domain.URLShortnerService
}

func NewHttpHandler(srvc domain.URLShortnerService) *HttpHandler {
	return &HttpHandler{service: srvc}
}

func (h *HttpHandler) ShortenURL(c *gin.Context) {
	req := ShortenURLRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	shortURL, err := h.service.ShortenURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error(), Message: "FAILED"})
		return
	}

	url := fmt.Sprintf("%s/v1/%s", c.Request.Host, shortURL)

	c.JSON(http.StatusOK, ShortenURLSuccessResponse{Message: "SUCCESS", ShortURL: url})
}

func (h *HttpHandler) RedirectToFullURL(c *gin.Context) {
	hash := c.Param("key")

	fullURL, err := h.service.GetOriginalURL(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error(), Message: "FAILED"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, fullURL)
}
