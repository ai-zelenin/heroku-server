package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	err := http.ListenAndServe(net.JoinHostPort("", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := handleBody(r.Body, func(body []byte) error {
			m := make(map[string]interface{})
			err := json.Unmarshal(body, &m)
			if err != nil {
				return err
			}
			data, err := json.MarshalIndent(m, "", "\t")
			if err != nil {
				return err
			}
			log.Println(string(data))
			return nil
		})
		if err != nil {
			http.Error(w, "handleBody error", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text/css")
		io.WriteString(w, `
.red {
    font-family: sans-serif;
    line-height: 1.15;
	color:red;
}`)
	}))
	if err != nil {
		log.Fatal(err)
	}
}

func handleBody(br io.ReadCloser, cb func(body []byte) error) error {
	body, err := ioutil.ReadAll(br)
	if err != nil {
		return err
	}
	defer br.Close()
	return cb(body)
}
