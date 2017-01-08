/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 07/01/2017    
* Time      : 16:10
* Developer : ibrahimcobani
*
*******/
package main

import (
	"net/http"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"fmt"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(ArithService), "")

	s.RegisterBeforeFunc(func(ri *rpc.RequestInfo) {
		fmt.Println(ri)
	})

	s.RegisterAfterFunc(func(ri *rpc.RequestInfo) {
		fmt.Println(ri)
	})

	http.Handle("/rpc", s)

	fmt.Println("http://localhost:1234 RPC server is starting")
	http.ListenAndServe(":1234", nil)
}

type Args struct {
	A, B int
}

type ArithService struct{}

func (this *ArithService) Add(req *http.Request, args *Args, reply *int) error {
	fmt.Println(args.A, args.B)
	*reply = args.A + args.B

	fmt.Println("Calculating")
	return nil
}