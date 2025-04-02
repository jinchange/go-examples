package main

import (
	"go-examples/gee"
	"log"
	"net/http"
)

func main() {
	StartHttpServer()
}

func StartHttpServer() {
	r := gee.New()
	r.GET("/hello", func(context *gee.Context) {
		context.JSON(http.StatusOK, gee.H{
			"name": context.Query("name"),
			"age":  context.Query("age"),
		})
	})
	log.Panic(r.Run(":9999"))
}
