// src: https://gobyexample.com/http-servers
package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// src: https://pkg.go.dev/net/http#ListenAndServeTLS
	// certificates can generate for local use by command:
	// openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out certificate.pem
	err := http.ListenAndServeTLS(":443", "certificate.pem", "key.pem", nil)
	log.Fatal(err)
}
