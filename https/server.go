package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server starting in localhost:8081....")
	if err := http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi jazz,this is an example of https service in golange!")
}
