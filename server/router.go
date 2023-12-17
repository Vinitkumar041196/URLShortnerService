package server

import (
	shortnerHTTPDelivery "url-shortner/url_shortner/delivery/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(srv *Server) *gin.Engine {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")

	handler := shortnerHTTPDelivery.NewHttpHandler(srv.urlShortnerService)
	v1.POST("/url/shorten", handler.ShortenURL)
	v1.GET("/url/:key", handler.RedirectToFullURL)
	return router
}
