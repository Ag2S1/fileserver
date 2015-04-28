package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Start file server in " + dir)

	handler := http.FileServer(http.Dir(dir))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})

	fmt.Println("listen to 127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
