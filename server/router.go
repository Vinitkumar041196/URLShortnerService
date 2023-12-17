package server

import (
	_ "url-shortener/docs"
	metricsHTTPDelivery "url-shortener/url_metrics/delivery/http"
	shortenerHTTPDelivery "url-shortener/url_shortener/delivery/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(srv *Server) *gin.Engine {
	router := gin.Default()
	//Swagger endpoint
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/v1")

	//Set up URL shortener routes
	NewURLShortenerRouter(v1, srv)

	//Set up metric routes
	NewMetricsRouter(v1, srv)
	return router
}

// Set up URL shortener routes
func NewURLShortenerRouter(v1 *gin.RouterGroup, srv *Server) {
	//Create a new URL shortener handler
	urlShortenerhandler := shortenerHTTPDelivery.NewURLShortenerHttpHandler(srv.urlShortenerService)
	//API to get shortened URL
	v1.POST("/url/shorten", urlShortenerhandler.ShortenURL)
	//API called for redirecting shortened URL to actual URL
	v1.GET("/:key", urlShortenerhandler.RedirectToFullURL)
}

// Set up metric routes
func NewMetricsRouter(v1 *gin.RouterGroup, srv *Server) {
	//Create a new URL shortener handler
	metricshandler := metricsHTTPDelivery.NewDomainMetricsHttpHandler(srv.metricsService)
	//API to get top N domains
	v1.GET("/metrics/domains/top", metricshandler.GetTopDomains)
}
