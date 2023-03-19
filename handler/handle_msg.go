package handler

import (
	"IM_component/model"
	"IM_component/server"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	Srv *server.Server
}

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	msgXml := model.TextMsgXML{}
	buf, _ := ioutil.ReadAll(r.Body)
	err := xml.Unmarshal(buf, &msgXml)
	if err != nil {
		return
	}

	Result := model.Result{
		Code:    1001,
		Message: "common",
	}
	Msg := model.Msg{
		Media_path: "",
		Content:    "你好呀",
	}

	Data := model.Data{
		Id:         msgXml.MsgId,
		Name:       "",
		Avatar:     "",
		Msg_type:   "",
		Contact_id: msgXml.ToUserName,
		Company_id: "",
		Send_Time:  time.Now().Minute(),
		Msg:        Msg,
	}

	TextMsg := model.TextMsg{
		Data:   Data,
		Result: Result,
	}

	j, err := json.Marshal(TextMsg)
	if err != nil {
		log.Fatal(err)
	}
	h.Srv.MsgBuffer[msgXml.ToUserName] = append(h.Srv.MsgBuffer[msgXml.ToUserName], j)
	w.Write([]byte(`{"result":{"Code":100,"Message":"成功"}}`))
}
