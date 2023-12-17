package server

import (
	_ "url-shortner/docs"
	metricsHTTPDelivery "url-shortner/url_metrics/delivery/http"
	shortnerHTTPDelivery "url-shortner/url_shortner/delivery/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(srv *Server) *gin.Engine {
	router := gin.Default()
	//Swagger endpoint
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/v1")

	//Set up URL shortner routes
	NewURLShortnerRouter(v1, srv)

	//Set up metric routes
	NewMetricsRouter(v1, srv)
	return router
}

// Set up URL shortner routes
func NewURLShortnerRouter(v1 *gin.RouterGroup, srv *Server) {
	//Create a new URL shortner handler
	urlShortnerhandler := shortnerHTTPDelivery.NewURLShortnerHttpHandler(srv.urlShortnerService)
	//API to get shortened URL
	v1.POST("/url/shorten", urlShortnerhandler.ShortenURL)
	//API called for redirecting shortened URL to actual URL
	v1.GET("/:key", urlShortnerhandler.RedirectToFullURL)
}

// Set up metric routes
func NewMetricsRouter(v1 *gin.RouterGroup, srv *Server) {
	//Create a new URL shortner handler
	metricshandler := metricsHTTPDelivery.NewDomainMetricsHttpHandler(srv.metricsService)
	//API to get top N domains
	v1.GET("/metrics/domains/top", metricshandler.GetTopDomains)
}
