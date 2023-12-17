package http

import (
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
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error(), Message: "FAILED"})
		return
	}

	c.JSON(http.StatusOK, ShortenURLSuccessResponse{Message: "SUCCESS", ShortURL: shortURL})
}
