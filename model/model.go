package model

type Result struct {
	Code    int8   `json:"code"`
	Message string `json:"message"`
}

type Msg struct {
	Media_path string `json:"media_path,omitempty"`
	Content    string `json:"content,omitempty"`
}

type Data struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Msg_type   string `json:"msg_type"`
	Contact_id string `json:"contact_id"`
	Company_id string `json:"company_id"`
	Send_Time  int    `json:"send_time"`
	Msg        Msg    `json:"msg"`
}

type SysMsg struct {
	Data   Data   `json:"data"`
	Result Result `json:"result"`
}

type TextMsg struct {
	Data   Data   `json:"data"`
	Result Result `json:"result"`
}
