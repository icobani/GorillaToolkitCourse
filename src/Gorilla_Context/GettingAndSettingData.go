/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 06/01/2017    
* Time      : 22:41
* Developer : ibrahimcobani
*
*******/
package Gorilla_Context

import (
	"net/http"
	"github.com/gorilla/context"
	"fmt"
)

const key = "MyKey"

func GettingAndSettingData() {

	request1 := &http.Request{}
	request2 := &http.Request{}


	//Set Value
	context.Set(request1, key, "foo")
	context.Set(request1, "second", "bidi")
	context.Set(request1, "other", "midi")
	context.Set(request2, key, "bar")


	values := context.GetAll(request1)

	fmt.Println("------")
	fmt.Println(values["second"])
	fmt.Println(values["other"])
	fmt.Println(values[key])
	fmt.Println("------")

	//Get Value
	fmt.Println(context.Get(request1, key))


	//Delete a Key
	context.Delete(request1,key)
	fmt.Println(context.Get(request1, key))


	//all delete values
	fmt.Println(context.Get(request2, key))
	context.Clear(request2)
	fmt.Println(context.Get(request2, key))

	context.pu
}