package main

import (
	"fmt"
	"geeExample/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,r.URL.Path)
	})
	r.POST("/hello",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,r.URL.Path)
	})
	r.Run(":9999")
}
