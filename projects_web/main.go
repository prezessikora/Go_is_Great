package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You requested %v", request.URL.Path)

}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
