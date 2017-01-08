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

//File base yaptığımız zaman sadece aşağıdaki kod değişiyor.
var filestore = sessions.NewFilesystemStore("src/tmp", securecookie.GenerateRandomKey(64))

func File_Base_Session() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, _ := filestore.Get(r, "session-name")
		session.Values["faworite-Icecream"] = []string{"vanilla", "butter pecan"}

		//Flash Mesaj ekleniyor.
		session.AddFlash("Flash message only work  once!")
		fmt.Println("Eklendi...")

		sessions.Save(r, w)
		w.Write([]byte("<h1>Hello File Base Session</h1><br><a href='/sessioncontents'>sessioncontents</a>"))
	})

	http.HandleFunc("/sessioncontents", func(w http.ResponseWriter, r *http.Request) {
		session, _ := filestore.Get(r, "session-name")
		favorites := session.Values["faworite-Icecream"].([]string)

		w.Header().Add("Content-Type", "text/html")


		//Flash Mesaj kontrolü..

		if flashes := session.Flashes(); len(flashes) > 0 {

			fmt.Println(len(flashes))

			session.Save(r, w)
			if flashes2 := session.Flashes(); len(flashes2) > 0 {
				fmt.Println(len(flashes2))
			}
			session.Save(r, w)


			for _, msg := range flashes {
				m := msg.(string)
				w.Write([]byte(m + " .... <br>"))
			}
		}

		w.Write([]byte("<hr><br>"))

		for _, flavor := range favorites {
			w.Write([]byte(flavor + "<br>"))
		}

		w.Write([]byte("Notting to see"))
		w.Write([]byte("<br><a href='/'>Go Back</a>"))
	})
	fmt.Println("http://localhost::3000 File Bese Session Store server is started.....")


	//Bunu yapmazsan memory leak yavaşça server'ı öldürür.....
	http.ListenAndServe(":3000", context.ClearHandler(http.DefaultServeMux))
}
