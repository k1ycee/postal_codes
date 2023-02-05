package main

import "net/http"

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// create response binary data
	data := []byte("This is a go server") // slice of bytes
	// write `data` to response
	res.Write(data)
}

func main() {
	// create a new handler
	handler := HttpHandler{}
	// listen and serve
	http.ListenAndServe(":3000", handler)
}
