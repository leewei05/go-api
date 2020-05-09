package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	serverPort = flag.String(
		"httpPort",
		os.Getenv("HTTP_PORT"),
		"Port number of HTTP server",
	)
)

func init() {

}

func main() {
	flag.Parse()
	if *serverPort == "" {
		log.Panic("Null HTTP port value")
	}

	port := fmt.Sprintf(":%v", *serverPort)
	log.Printf("HTTP server running on port %v", port)
	http.ListenAndServe(port, nil)
}
