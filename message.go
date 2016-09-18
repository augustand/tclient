package main

type Message struct {
	MsgType string      // 消息类型,text,json
	Event   string      // 事件名字
	MsgBody interface{} // 消息体
}
