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
	"html/template"
	"os"
)

func RouteTemplate() {
	regexp, _ := reverse.CompileRegexp(`/foo/(?P<bar>.+)`)

	reverter := func(regexp *reverse.Regexp) func(... string) (string, error) {
		return func(params ... string) (string, error) {
			values := url.Values{}
			for i := 0; i < len(params); i += 2 {
				values[params[i]] = []string{params[i + 1]}
			}
			return regexp.Revert(values)
		}
	}

	pathTemplete, _ := regexp.Revert(url.Values{"bar":{"bar:[0-9]+"}})
	path, _ := regexp.Revert(url.Values{"bar":{"42"}})

	fm := template.FuncMap{"pathGen": reverter(regexp)}

	templateRaw := `{{ pathGen "bar" "42" }}`

	t, _ := template.New("").Funcs(fm).Parse(templateRaw)

	fmt.Println("")
	fmt.Println(pathTemplete)
	fmt.Println(path)

	t.Execute(os.Stdout, nil)
	fmt.Println("")
	fmt.Println("")
}
