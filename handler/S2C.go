package handler

import (
	"IM_component/model"
	"context"
	"encoding/json"
	"log"

	"nhooyr.io/websocket"
)

// 处理服务端向客户端发消息
type S2C struct {
	conn *websocket.Conn
}

func (c S2C) SendSysMsg(m *model.SysMsg) {
	text, err := json.Marshal(m)
	if err != nil {
		log.Printf("%v", err)
	}
	c.conn.Write(context.Background(), websocket.MessageText, text)
}

func (c S2C) SendMsg(m *model.TextMsg) {

}
