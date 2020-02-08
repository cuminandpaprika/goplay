package server

import (
	"fmt"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", handler)
	// If dev, localhost
	host := "localhost"
	port := "8080"
	address := host + ":" + port
	fmt.Println("Serving at" + address)
	fmt.Println(http.ListenAndServe(address, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Here we get a request and we need to write the response
	fmt.Fprintf(w, "Hi there, I love coffee!☕")
}
