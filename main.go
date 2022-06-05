package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/pages.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
func main() {
	fmt.Printf("Hello %s/%s\n", runtime.GOOS, runtime.GOARCH)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
