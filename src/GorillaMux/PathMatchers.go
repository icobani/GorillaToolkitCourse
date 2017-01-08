/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/01/2017    
* Time      : 19:26
* Developer : ibrahimcobani
*
*******/
package GorillaMux

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func PathMatchers() {

	r := mux.NewRouter()

	func1 := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		w.Write([]byte(vars["id"] + " - " + vars["name"]))
	}

	r.HandleFunc("/foo/{id:[0-9]+}/{name}", func1)

	http.Handle("/", r)

	go http.ListenAndServe(":4000", nil)

	fmt.Println("http://localhost:4000/foo/123 ")
	fmt.Println("Server is running now")
	fmt.Scanln()
}
