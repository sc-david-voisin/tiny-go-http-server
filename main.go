package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func isHTTPS(req *http.Request) bool {
  return req.Header.Get("X-Forwarded-Proto") == "https" || req.URL.Scheme == "https"
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

  log.Fatalln(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}

