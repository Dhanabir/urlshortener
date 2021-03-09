package main

import (
	"fmt"
	"net/http"
	"log"
)

func init() {
	err := LoadFile("urls.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func handleURL(w http.ResponseWriter, r *http.Request) {
	longUrl := r.URL.Path[1:]
	fmt.Fprint(w, SaveAndSendUrl(longUrl))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handleURL)
	log.Fatal(http.ListenAndServe(":80", nil))
}