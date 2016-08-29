package main

import (
	"runtime"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := NewClient("127.0.0.1", "1235")
	if err != nil {
		println(err)
	}
	defer c.Close()

	c.OnConnect(func(c *Client) {
		fmt.Println("on connect")
	})

	c.OnMessage(func(d string) {
		fmt.Println(d)
	})

	c.OnDisConnect(func(c *Client) {
		fmt.Println("on disconnect")
	})

}





