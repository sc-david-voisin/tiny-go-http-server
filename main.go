package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func isHTTPS(req *http.Request) bool {
	return req.Header.Get("X-Forwarded-Proto") == "https" || req.URL.Scheme == "https"
}

func ensurePort(defaultPort uint16) string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return port
	}
	return strconv.FormatUint(uint64(defaultPort), 10)
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if isHTTPS(req) {
			log.Println("HTTPS is used, connection is secured.")
		} else {
			log.Println("HTTP is used, you should envisage using HTTPS.")
		}
		fmt.Fprintf(res, "Hello World\n")
	})

	port := ensurePort(6666)
	log.Printf("Listen on port %s\n", port)

	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
