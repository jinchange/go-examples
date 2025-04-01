package main

import (
	"encoding/json"
	"fmt"
	"go-examples/gee"
	"net/http"
)

func main() {
	StartHttpServer()
}

func StartHttpServer() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.GET("/user", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		user := map[string]interface{}{
			"name":     "geektutu",
			"password": "1234",
		}
		encoder := json.NewEncoder(writer)
		if err := encoder.Encode(user); err != nil {
			http.Error(writer, err.Error(), 500)
		}
	})

	r.Run(":9999")
}
