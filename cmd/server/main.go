package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("cj")
	}()

	serverHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("meheheheheh", "goohehehe")
		w.WriteHeader(404)
		for _, id := range req.URL.Query()["id"] {
			_, _ = io.WriteString(w, id)
		}
	}

	http.HandleFunc("/issues", serverHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
