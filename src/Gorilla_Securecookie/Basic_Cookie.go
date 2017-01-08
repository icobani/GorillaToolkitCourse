/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 07/01/2017    
* Time      : 15:02
* Developer : ibrahimcobani
*
*******/
package Gorilla_Securecookie

import (
	"net/http"
	"fmt"
	"github.com/gorilla/securecookie"
)

func Basic_Cookie() {

	hashKey := securecookie.GenerateRandomKey(64)
	blockKey := securecookie.GenerateRandomKey(32)

	s := securecookie.New(hashKey, blockKey)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if encoded, err := s.Encode("myCookie", "myvalue"); err == nil {

			cookie := http.Cookie{
				Name: "mycookie",
				Value:encoded,
				Path:"/",
			}

			var result string
			if err = s.Decode("myCookie", cookie.Value, &result); err == nil {
				fmt.Println(" ==> " + result)
			}

			http.SetCookie(w, &cookie)
		}
		w.Write([]byte("Setting cookies"))
	})
	fmt.Println("http://localhost::3000")
	http.ListenAndServe(":3000", nil)

}
