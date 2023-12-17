package server

import (
	shortnerHTTPDelivery "url-shortner/url_shortner/delivery/http"
	_ "url-shortner/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(srv *Server) *gin.Engine {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/v1")

	//Create a new URL shortner handler
	urlShortnerhandler := shortnerHTTPDelivery.NewHttpHandler(srv.urlShortnerService)
	//API to get shortened URL
	v1.POST("/url/shorten", urlShortnerhandler.ShortenURL)
	//API called for redirecting shortened URL to actual URL 
	v1.GET("/:key", urlShortnerhandler.RedirectToFullURL)
	return router
}
