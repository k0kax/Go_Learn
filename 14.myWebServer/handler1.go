package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

// 自定义handler
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,world"))
}

func main() {
	mh := MyHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}
	fmt.Println(">>>>>>>>>>>>>web server working>>>>>>>>")
	server.ListenAndServe()

}
