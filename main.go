package main

import (
	"IM_component/handler"
	"IM_component/server"
	"log"
	"net/http"
)

func main() {
	sv := server.Server{KeepAliving: false, MsgBuffer: make(map[string][][]byte)}
	http.HandleFunc("/", sv.Handler)

	h := handler.Handler{Srv: &sv}
	http.HandleFunc("/send", h.Send)

	err := http.ListenAndServe("localhost:1234", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
