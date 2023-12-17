package main

import (
	"log"
	"url-shortner/server"
)

func main() {
	//create a new HTTP server
	srv := server.NewServer()
	//Start HTTP Server
	srv.Start()
	log.Println("HTTP Server Exited")
}
