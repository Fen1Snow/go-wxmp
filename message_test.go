package wxmp

import (
	"encoding/xml"
	"net/http"
	"testing"
)

func TestMessageText(t *testing.T) {
	data := `
		<xml>
		  <ToUserName><![CDATA[toUser]]></ToUserName>
		  <FromUserName><![CDATA[fromUser]]></FromUserName>
		  <CreateTime>1348831860</CreateTime>
		  <MsgType><![CDATA[text]]></MsgType>
		  <Content><![CDATA[this is a test]]></Content>
		  <MsgId>1234567890123456</MsgId>
		</xml>
	`
	r := MsgText{}
	err := xml.Unmarshal([]byte(data), &r)
	if err != nil {
		panic(err)
	}
}

func TestMessage_HttpServer(t *testing.T) {
	http.HandleFunc("/test", c.Message().HttpServer())
	http.ListenAndServe(":40018", nil)
}
