package main

import (
	"fmt"
	"net/http"
)

// the handler of http request and response
func handler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprint(writer, "Hello world!", request.URL.Path[1:])
}


// the main function of project.
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
