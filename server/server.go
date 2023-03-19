package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct {
	ConnMax     int
	KeepAliving bool
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: false,
	Subprotocols:      []string{"echo"},
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (e *Server) Handler(w http.ResponseWriter, r *http.Request) {

	if e.ConnMax+1 > 10 {
		// w.Write()
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	// defer conn.Close()
	// 连接建立失败
	if err != nil {
		log.Printf("%v", err)
		return
	}

	// websocket 连接成功，进入websocket持续监听
	// 心跳
	go Ping(conn)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	for {
		<-time.After(3 * time.Second)
		msg, err := GetMsg(conn)
		if err != nil {
			conn.Close()
			break
		}
		err = SendMsg(conn, msg)
		if err != nil {
			conn.Close()
			break
		}
		fmt.Printf("%s", msg)
	}
}

func Ping(conn *websocket.Conn) {

	for {
		// 每1s发一次心跳，保活
		<-time.After(1 * time.Second)

		err := conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			fmt.Println("连接已断开")
			conn.Close()
			break
		} else {
			// 连接再保持5秒
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		}
	}
}

func GetMsg(conn *websocket.Conn) (string, error) {
	fmt.Println("getmsg：")

	_, p, err := conn.ReadMessage()
	if err != nil {
		return "", err
	}
	// fmt.Printf("%s", string(p))
	return string(p), nil
}

func SendMsg(conn *websocket.Conn, message string) error {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
