package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"url-shortner/domain"
)

type Server struct {
	urlShortnerService *domain.URLShortnerService
}

func NewServer() *Server {
	return &Server{}
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
