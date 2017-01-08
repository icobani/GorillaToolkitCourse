/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/01/2017    
* Time      : 22:22
* Developer : ibrahimcobani
*
*******/
package GorillaReverse

import (
	"net/url"
	"net/http"
	"github.com/gorilla/reverse"
	"fmt"
)

func BasicMatchers() {
	u := &url.URL{
		Scheme:"http",
		Host:"localhost:9999",
		Path:"/foo/55",
		RawQuery:"buz=ss",
	}

	r := &http.Request{URL:u}

	p, _ := reverse.NewRegexpPath("/foo/[0-9]+")
	fmt.Println("Match?: ", p.Match(r))
}