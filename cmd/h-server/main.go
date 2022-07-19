package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	err := http.ListenAndServe(net.JoinHostPort("", port), LogMiddleware(http.FileServer(http.Dir("./static"))))
	if err != nil {
		log.Fatal(err)
	}
}
func LogMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("addr:%s url:%s", r.RemoteAddr, r.URL)
		handler.ServeHTTP(w, r)
	})
}
func handleBody(br io.ReadCloser, cb func(body []byte) error) error {
	body, err := ioutil.ReadAll(br)
	if err != nil {
		return err
	}
	defer br.Close()
	return cb(body)
}
