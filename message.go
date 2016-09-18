package main

type Message struct {
	Order string `json:"order,omitempty"` // 指令名字
	Data  string `json:"data,omitempty"`  // 消息体
}


