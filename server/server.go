package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"url-shortner/domain"
	metricsRepo "url-shortner/url_metrics/repository"
	metricsService "url-shortner/url_metrics/service"
	urlRepo "url-shortner/url_shortner/repository"
	urlShortnerService "url-shortner/url_shortner/service"
)

type Server struct {
	urlShortnerService domain.URLShortnerService
	metricsService     domain.DomainMetricsService
}

func NewServer() *Server {
	//initialize metrics store
	metricsStore := metricsRepo.NewInMemoryMetricStore()
	//setting up metrics service
	metricsSrvc := metricsService.NewDomainMetricsService(metricsStore)

	//setting up url shortner service
	urlShortnerSrvc := urlShortnerService.NewURLShortnerService(urlRepo.NewInMemoryURLStore(), metricsStore)

	return &Server{urlShortnerService: urlShortnerSrvc, metricsService: metricsSrvc}
}

func (srv *Server) Start() {
	router := NewRouter(srv)

	httpServer := &http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
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
