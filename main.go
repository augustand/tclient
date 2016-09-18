package main

import (
	"runtime"
	"fmt"
	"encoding/json"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := NewClient("127.0.0.1", "1235")
	defer c.Close()

	c.OnConnect(func() {
		fmt.Println("连接成功")
		c.Sent("connect")
	})

	c.OnMessage(func(d string) {
		fmt.Println(c.Conn.LocalAddr().String())
		fmt.Println(d)

		m := &Message{}
		err := json.Unmarshal([]byte(d), m)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(m.Data)
		fmt.Println(m.Order)

	})

	c.OnDisConnect(func() {
		fmt.Println("on disconnect")
	})

}





