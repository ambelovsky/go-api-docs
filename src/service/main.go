package main

import (
	"fmt"
	"net/http"
	"service/internal/apidoc"
)

func main() {
	http.HandleFunc("/", apidoc.HTTPHandlerWithPrefix("/"))
	//...
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
