package main

import (
	"strings"
	"io"
	"github.com/sadlil/go-trigger"
	"net"
	"bufio"
	"time"
	"fmt"
	"math"
)

type Client struct {
	Name          string            // 设置客户端的名字,默认为ip:port
	Address       string            // 服务地址
	Port          string            // 服务端口
	conn          chan *net.TCPConn // 当前的连接，如果 nil 表示没有连接
	Conn          *net.TCPConn
	maxRetry      int               // 最大重试次数
	quit          chan bool         // 当服务器发出结束的标志时退出
	isAlive       bool              // 判断是否连接成功
	ConnTime      time.Time         // 连接时间
	VerifyKey     string            // 连接验证KEY
	ConnVerify    bool              // 是否验证
	isExceptional chan bool         // 连接发生异常
}

// 连接成功后回调这个函数
func (c *Client)OnConnect(f func()) {
	trigger.On(OnConnect, f)
	if c.Conn != nil {
		f()
	}
}

func (c *Client)reTry() {
	if c.maxRetry == 6 {
		c.maxRetry = 1
		return
	}

	t := time.Duration(math.Pow(2, float64(c.maxRetry)))
	fmt.Println("等待时间%d", t)
	time.Sleep(t * time.Second)
	c.maxRetry += 1
}

func (c *Client)OnDisConnect(f func()) {
	trigger.On(OnDisConnect, f)
}

func (c *Client)OnError(f func(error)) {
	trigger.On(OnError, f)
}

func (c *Client) Close() {
	if <-c.quit && c.Conn != nil {
		fmt.Println("do quit")
		c.Conn.Close()

	}
}

func (c *Client) Sent(data string) {
	go func(c *Client) {
		_, err := c.Conn.Write([]byte(data))
		if err != nil {
			println(err.Error())
		}
		//println(n)
	}(c)
}

func (c *Client) OnMessage(f func(data string)) {

	go func(c *Client) {
		for {
			select {
			case conn := <-c.conn:

			//c.Conn.SetReadDeadline(time.Now().Add(time.Second * 4))
				time.Sleep(time.Millisecond * 100)

			//println("接收消息")
				reader := bufio.NewReader(conn)
				try1:
				for {
					msg, err := reader.ReadString('\n')
					msg = strings.Trim(msg, "\r\n")
					if err == io.EOF {
						println("暂时没有消息")
						c.isExceptional <- true
						trigger.Fire(OnDisConnect)
						break try1
					} else {
						f(msg)
					}
				}
			//println("结束消息")
			}
		}
	}(c)
}



// 重新连接服务器
func (c *Client)reConnect() {

	go func(c *Client) {
		for {
			select {
			case <-c.isExceptional:

				var (
					err error
					conn *net.TCPConn
				)

				conn, err = connect(c.Address, c.Port)
				if err != nil {
					c.isExceptional <- true
					c.Conn = nil
					fmt.Println("尝试重新连接")
					c.reTry()
				} else {
					c.isAlive = true
					c.maxRetry = 1
					c.Conn = conn
					c.conn <- conn
					c.isExceptional = make(chan bool, 100)

					trigger.Fire(OnConnect)
					//fmt.Println("连接成功")
				}
			}
		}
	}(c)
}

// 连接服务器
func connect(addr, port string) (*net.TCPConn, error) {
	var (
		err error
		tcpAddr *net.TCPAddr
		conn *net.TCPConn
	)

	tcpAddr, err = net.ResolveTCPAddr("tcp", addr + ":" + port) //获取一个TCP地址信息,TCPAddr
	if err != nil {
		return nil, err
	}

	conn, err = net.DialTCP("tcp", nil, tcpAddr) //创建一个TCP连接:TCPConn
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewClient(addr, port string) (*Client) {
	c := Client{}
	c.maxRetry = 1

	c.quit = make(chan bool)
	c.isExceptional = make(chan bool, 100)
	c.conn = make(chan *net.TCPConn, 2)

	c.Address = addr
	c.Port = port
	c.isExceptional <- true
	c.reConnect()
	return &c
}


//else if err != nil {

//a := c.isconnect()
//
//if a != nil {
//	println("断开了连接")
//	c.Conn.Close()
//	c.isExceptional <- true
//	println(len(c.isExceptional))
//	println("isExceptional")
//	println(a.Error())
//	break try1
//}

//if strings.Index(err.Error(), "timeout") != -1 {
//	println(len(c.isExceptional))
//	println("isExceptional")
//	println(err.Error())
//	time.Sleep(time.Second * 5)
//	c.isExceptional <- true
//	break try1
//}

//a := c.connect()
//if a != nil {
//c.isExceptional <- true
//println("isExceptional")
//println(err.Error())
//	println(len(c.isExceptional))
//	println(a.Error())
//	println("isExceptional")
//break try1
//}
//
//println(err.Error())
//time.Sleep(time.Second * 2)
//println(err.Error())
//break try1

//c.reConnect()
//c.quit <- true
//}

//go func() {
//	for {
//
//		select {
//		//case <-c.isOK:
//		//	println("ok")
//		case <-time.After(10 * time.Second):
//
//			if c.isconnect() != nil {
//				c.isExceptional <- true
//			}
//		}
//	}
//}()


//else {
//		conn.SetDeadline(time.Now().Add(time.Second * 2))
//		conn.SetKeepAlive(true)
//	}