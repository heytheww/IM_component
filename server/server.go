package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct {
	KeepAliving bool
	MsgBuffer   map[string][][]byte // 字节数组的数组
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

	conn, err := upgrader.Upgrade(w, r, nil)

	// 连接建立失败
	if err != nil {
		log.Printf("%v", err)
		return
	}

	// websocket 连接成功
	userId := r.URL.Query().Get("userId")

	// 进入websocket持续监听
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	for {
		<-time.After(5 * time.Second)

		go func() {
			msg, err := GetMsg(conn)
			if err != nil {
				conn.Close()
			}
			fmt.Printf("%s", msg)
		}()

		if len(e.MsgBuffer[userId]) > 0 {
			for _, v := range e.MsgBuffer[userId] {
				err = SendMsg(conn, v)
				if err != nil {
					conn.Close()
				}
			}

			e.MsgBuffer[userId] = make([][]byte, 0)
		}

		// 心跳保活
		err := conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			fmt.Println("连接已断开")
			conn.Close()
			break
		} else {
			// 连接再保持10秒
			conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		}
	}
}

func Ping(conn *websocket.Conn, ch chan int) {

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

func SendMsg(conn *websocket.Conn, message []byte) error {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
