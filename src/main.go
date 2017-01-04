package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	func1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("func 1"))
	}

	func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("func 2"))
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.TLS == nil {
			func1(w, r)
		} else {
			func2(w, r)
		}
	}

	r.HandleFunc("/", handler)

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)
	go http.ListenAndServeTLS(":4443", "cert.pem", "key.pem", nil)

	fmt.Println("http://localhost:4000 ")
	fmt.Println("Server is running now")
	fmt.Scanln()
}