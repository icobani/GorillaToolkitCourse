/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 07/01/2017    
* Time      : 12:56
* Developer : ibrahimcobani
*
*******/
package Gorilla_Schema

import (
	"fmt"
	"github.com/gorilla/schema"
)

func Basic_Form_Decoding() {

	form := map[string][]string{
		"firstName":[]string{"ibrahim"},
		"lastName":[]string{"COBANİ"},
		"addresses.0.city":[]string{"Istanbul"},
		"addresses.0.state":[]string{"Fatih"},
		"addresses.1.city":[]string{"Elazığ"},
		"addresses.1.state":[]string{"Merkez"},
	}

	var P Person

	d := schema.NewDecoder()
	d.Decode(&P, form)

	fmt.Println(P)
}

type Person struct {
	FirstName string `schema:"firstName"`
	LastName  string `schema:"lastName"`
	Addresses []Address `schema:"addresses"`
}

type Address struct {
	City  string `schema:"city"`
	State string `schema:"state"`
}