package model

import "encoding/xml"

type TextMsgXML struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        int      `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
}
