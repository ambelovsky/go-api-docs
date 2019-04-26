package main

import (
	"fmt"
	"net/http"
	"service/internal/apidoc"
)

func main() {
	http.HandleFunc("/help/", apidoc.HTTPHandlerWithPrefix("/help"))
	//...
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
