package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

const validTimeFormat = "20060102T150405Z"

func main() {
	address := flag.String("address", "127.0.0.1", "the address of the http httpServer")
	port := flag.Int("port", 8080, "the port the http httpServer listens on")
	endpoint := flag.String("endpoint", "/ptlist", "the exposed endpoint")
	supportedPeriods := flag.String("periods", "1h,1d,1mo,1y", "the supported time periods")
	flag.Parse()

	//It is assumed that values provided when setting up the server are valid and properly formatted
	validator := newValidator(*endpoint, validTimeFormat, strings.Split(*supportedPeriods, ","))

	log.Printf("Start server listening at %s:%d", *address, *port)
	server := newHttpServer(*address, *port, validator)
	log.Fatal(http.ListenAndServe(server.getAddressAndPort(), server))
}
