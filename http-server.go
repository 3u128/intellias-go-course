// src: https://gobyexample.com/http-servers
package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// src: https://pkg.go.dev/net/http#ListenAndServeTLS
	// certificates can generate for local use by command:
	// openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out certificate.pem
	err := http.ListenAndServeTLS(":443", "certificate.pem", "key.pem", nil)
	log.Fatal(err)
}
