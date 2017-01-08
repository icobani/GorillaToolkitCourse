/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/01/2017    
* Time      : 19:06
* Developer : ibrahimcobani
*
*******/

package GorillaMux

import (
"github.com/gorilla/mux"
"fmt"
"net/http"
)

func PathPrefixMatchers() {

	r := mux.NewRouter()

	func1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("func 1..."))
	}

	func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("func 2..."))
	}


	r.PathPrefix("/foo").HandlerFunc(func1)
	r.PathPrefix("/bar").HandlerFunc(func2)

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)

	fmt.Println("http://localhost:4000 ")
	fmt.Println("Server is running now")
	fmt.Scanln()
}