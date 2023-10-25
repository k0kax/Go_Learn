package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

// 自定义handler
func (m *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,world"))
}

type AboutHandler struct{}

// 自定义handler
func (m *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about"))
}
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func main() {
	hello := HelloHandler{}
	about := AboutHandler{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	fmt.Println(">>>>>>>>>>>>>web server working>>>>>>>>")
	http.Handle("/hello", &hello) //自定义的handle
	http.Handle("/about", &about)

	//HandlerFunc("路径",方法)
	http.HandleFunc("/home", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("home"))
	})

	//http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()

}
