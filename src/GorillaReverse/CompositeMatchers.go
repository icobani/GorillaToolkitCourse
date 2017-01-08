/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 05/01/2017    
* Time      : 22:51
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

func CompositeMatchers() {
	u := &url.URL{
		Scheme:"http",
		Host:"localhost:9999",
		Path:"/foo/55",
		RawQuery:"buz=78",
	}

	r := &http.Request{URL:u}

	p, _ := reverse.NewRegexpPath("/foo/[0-9]+")
	q := reverse.NewQuery(map[string]string{"buz": "78"})
	//a := reverse.NewAll([]reverse.Matcher{p, q})
	//n := reverse.NewNone()
	o := reverse.NewOne([]reverse.Matcher{p, q})

	fmt.Println("Match?: ", o.Match(r))
}