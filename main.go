package main

import (
	"IM_component/server"
	"log"
	"net/http"
)

func main() {
	sv := server.Server{ConnMax: 10, KeepAliving: false}
	http.HandleFunc("/", sv.Handler)

	err := http.ListenAndServe("localhost:1234", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
