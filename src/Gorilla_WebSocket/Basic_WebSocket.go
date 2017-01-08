/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 08/01/2017    
* Time      : 14:17
* Developer : ibrahimcobani
*
*******/
package Gorilla_WebSocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"io"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize    : 1024,
	WriteBufferSize    :1024,
}

func Basic_WebSocket() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Console aşağıdaki kod yazılır
		/*
			var ws = new WebSocket("ws://localhost:3000/ws");
			ws.addEventListener("message",function(e){console.log(e);})
			ws.send("Noliii la");
		*/
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		go func() {
			for {
				messagetype, reader, _ := conn.NextReader()

				writer, _ := conn.NextWriter(messagetype)

				io.Copy(writer, reader)
				writer.Close()
			}
		}()

	})

	http.ListenAndServe(":3000", nil)
}
