/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/01/2017    
* Time      : 19:52
* Developer : ibrahimcobani
*
*******/
package GorillaMux

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func QueryMatchers() {

	r := mux.NewRouter()

	func1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("func 1..."))
	}

	func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("func 2..."))
	}

	r.Path("/").Queries("foo", "bar").HandlerFunc(func1)
	r.Path("/").Queries("bar", "{bar:[0-9]+}").HandlerFunc(func2)

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)

	fmt.Println("http://localhost:4000/?bar=12 ")
	fmt.Println("http://localhost:4000/?foo=bar ")
	fmt.Println("Server is running now")
	fmt.Scanln()
}

