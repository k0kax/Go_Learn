package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,world"))
	})

	http.ListenAndServe("localhost:8080", nil) //创建HTTP，使用DefaultServeMux
}
