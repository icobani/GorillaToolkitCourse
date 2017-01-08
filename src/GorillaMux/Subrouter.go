/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/01/2017    
* Time      : 20:53
* Developer : ibrahimcobani
*
*******/
package GorillaMux

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func Subrouter() {

	r := mux.NewRouter()

	func1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("func 1..."))
	}

	func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("func 2..."))
	}

	http.HandleFunc("/", func1)
	sr := r.PathPrefix("/foo").Subrouter()
	sr.HandleFunc("/", func2)

	go http.ListenAndServe(":4000", nil)

	fmt.Println("http://localhost:4000/?bar=12 ")
	fmt.Println("http://localhost:4000/?foo=bar ")
	fmt.Println("Server is running now")
	fmt.Scanln()
}
