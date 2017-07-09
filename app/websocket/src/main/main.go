package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"io"
)

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	// configure log format
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	// handle websocket requests
	http.HandleFunc("/ws",
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(echoHandler)}
			s.ServeHTTP(w, req)
		})
	// handle static file requests
	http.Handle("/", http.FileServer(http.Dir("static")))

	// listen from port 8100
	log.Fatal(http.ListenAndServe(":8100",nil))
}
