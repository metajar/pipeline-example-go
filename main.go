package main

import (
	"fmt"
	"log"
	"net/http"
)

const webContent = "<h1>Hello World!</h1>"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
