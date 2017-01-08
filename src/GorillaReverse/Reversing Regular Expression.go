/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 06/01/2017    
* Time      : 19:59
* Developer : ibrahimcobani
*
*******/
package GorillaReverse

import (
	"github.com/gorilla/reverse"
	"net/url"
	"fmt"
)

func ReversingRegularExpression() {
	regexp, _ := reverse.CompileRegexp(`/foo/(?P<bar>\d+)`)
	r, _ := regexp.Revert(url.Values{"bar":{"89"}})

	fmt.Println(r)
}
