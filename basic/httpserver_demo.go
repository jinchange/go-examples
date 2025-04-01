package basic

import (
	"fmt"
	"log"
	"net/http"
)

// StartSimpleHttpServer basic demo
func StartSimpleHttpServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

type HttpServerHandler struct {
}

func (httpServer *HttpServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/hello" {
		fmt.Fprintf(w, "hello world %s\n", path)
	} else {
		fmt.Fprintln(w, path+" not find")
	}
}

func StartCustomHttpServer() {
	httpserver := new(HttpServerHandler)
	// 传入一个结构体，结构体实现了ServeHTTP接口，定义路由跳转逻辑
	log.Fatal(http.ListenAndServe("127.0.0.1:8888", httpserver))
}
