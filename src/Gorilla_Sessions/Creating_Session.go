/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 07/01/2017    
* Time      : 17:58
* Developer : ibrahimcobani
*
*******/
package Gorilla_Sessions

import (
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/context"
	"fmt"
)

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))

func Creating_Session() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		session.Values["faworite-Icecream"] = []string{"vanilla", "butter pecan"}
		sessions.Save(r, w)
		w.Write([]byte("<h1>Hello Session</h1>"))
	})

	http.HandleFunc("/sessioncontents", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		favorites := session.Values["faworite-Icecream"].([]string)
		w.Header().Add("Content-Type","text/html")

		for _, flavor := range favorites  {
			w.Write([]byte(flavor + "<br>"))
		}
		
		w.Write([]byte("Notting to see"))
	})
	fmt.Println("http://localhost::3000 server is started.....")


	//Bunu yapmazsan memory leak yavaşça server'ı öldürür.....
	http.ListenAndServe(":3000", context.ClearHandler(http.DefaultServeMux))
}
