package main

import (
	_ "embed"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//go:embed gif.gif
var gif []byte

//go:embed png.png
var png []byte

func main() {
	http.HandleFunc("/image.png", func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		log.Printf("%v %v: accept %v\n", r.Method, r.URL, accept)
		if strings.Contains(accept, "webp") {
			w.Header().Add("content-type", "image/gif")
			w.Header().Add("content-length", strconv.Itoa(len(gif)))
			w.Write(gif)
		} else {
			w.Header().Add("content-type", "image/png")
			w.Header().Add("content-length", strconv.Itoa(len(png)))
			w.Write(png)
		}
	})

	// diff with above by vary header
	http.HandleFunc("/image_with_vary.png", func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		log.Printf("%v %v: accept %v\n", r.Method, r.URL, accept)
		if strings.Contains(accept, "webp") {
			w.Header().Add("content-type", "image/gif")
			w.Header().Add("content-length", strconv.Itoa(len(gif)))
			w.Header().Add("vary", "accept")
			w.Write(gif)
		} else {
			w.Header().Add("content-type", "image/png")
			w.Header().Add("content-length", strconv.Itoa(len(png)))
			w.Write(png)
		}
	})
	log.Println("listen on :1236")
	http.ListenAndServe(":1236", nil)
}
