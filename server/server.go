package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"url-shortener/domain"
	metricsRepo "url-shortener/url_metrics/repository"
	metricsService "url-shortener/url_metrics/service"
	urlRepo "url-shortener/url_shortener/repository"
	urlShortenerService "url-shortener/url_shortener/service"
)

type Server struct {
	urlShortenerService domain.URLShortenerService
	metricsService      domain.DomainMetricsService
}

func NewServer() *Server {
	//initialize metrics store
	metricsStore := metricsRepo.NewInMemoryMetricStore()
	//setting up metrics service
	metricsSrvc := metricsService.NewDomainMetricsService(metricsStore)

	//setting up url shortener service
	urlShortenerSrvc := urlShortenerService.NewURLShortenerService(urlRepo.NewInMemoryURLStore(), metricsStore)

	return &Server{urlShortenerService: urlShortenerSrvc, metricsService: metricsSrvc}
}

func (srv *Server) Start() {
	router := NewRouter(srv)
	serverAddr := os.Getenv("SERVER_ADDR")
	if os.Getenv("SERVER_ADDR") == "" {
		serverAddr = ":8080"
	}
	httpServer := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	log.Println("Starting HTTP Server on ", httpServer.Addr)

	if os.Getenv("ENABLE_TLS") == "true" {
		log.Fatal(httpServer.ListenAndServeTLS(os.Getenv("SSL_CRT_PATH"), os.Getenv("SSL_KEY_PATH")))
	} else {
		//Globally disabling SSL certificate check
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		log.Fatal(httpServer.ListenAndServe())
	}
}
