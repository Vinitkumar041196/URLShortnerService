package main

import (
	"log"
	"url-shortner/server"
)

// @title URL SHORTNER
// @version 1.0
// @BasePath /v1
func main() {
	//create a new HTTP server
	srv := server.NewServer()
	//Start HTTP Server
	srv.Start()
	log.Println("HTTP Server Exited")
}
