package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "14.myWebServer/wwwroot"+r.URL.Path) //此处路径写的有些问题
		//http.ServeFile(w, r, "wwwroot/index.html")
	})
	http.ListenAndServe(":8080", nil)
}
