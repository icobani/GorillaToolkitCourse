/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/01/2017    
* Time      : 19:39
* Developer : ibrahimcobani
*
*******/
package GorillaMux

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func MethodMatchers() {

	r := mux.NewRouter()

	html := `
		<html><head></head><body>
		<form action = '' method='POST'>
			The field <input name = 'field'/>
		</form>
		</body></html>
	`

	func1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	}

	func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("field")))
	}

	r.HandleFunc("/", func1).Methods("GET")
	r.HandleFunc("/", func2).Methods("POST")

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)

	fmt.Println("http://localhost:4000/foo/123 ")
	fmt.Println("Server is running now")
	fmt.Scanln()
}
