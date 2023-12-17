package main

import (
	"log"
	"url-shortener/server"
)

// @title URL SHORTENER
// @version 1.0
// @BasePath /v1
func main() {
	//create a new HTTP server
	srv := server.NewServer()
	//Start HTTP Server
	srv.Start()
	log.Println("HTTP Server Exited")
}
