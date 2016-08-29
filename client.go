package main

import (
	"strings"
	"io"
	"github.com/sadlil/go-trigger"
	"net"
	"bufio"
	"time"
)

type Client struct {
	Address    string       // 服务地址
	Port       string       // 服务端口
	Conn       *net.TCPConn // 当前的连接，如果 nil 表示没有连接
	maxRetry   int          // 最大重试次数
	quit       chan bool    //当服务器发出结束的标志时退出
	isAlive    bool         // 判断是否连接成功
	ConnTime   time.Time    //连接时间
	VerifyKey  string       //连接验证KEY
	ConnVerify bool         //是否验证
}

func (c *Client)OnConnect(f func(c *Client)) {
	if c.isAlive {
		f(c)
	}
}

func (c *Client)OnDisConnect(f func(c *Client)) {
	trigger.On(OnDisConnect, f)
}

func (c *Client)OnError(f func(error)) {
	trigger.On(OnError, f)
}

func (c *Client) Close() {
	quit := <-c.quit
	if quit {
		c.Conn.Close()
	}

}

func (c *Client) OnMessage(f func(data string)) {
	go (func(conn *net.TCPConn) {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			msg = strings.Trim(msg, "\r\n")
			if err == io.EOF {
				trigger.Fire(OnDisConnect, c)
				c.quit <- true
				break
			} else if err != nil {
				trigger.Fire(OnError, err)
				c.quit <- true
				break
			} else {
				f(msg)
			}
		}
	})(c.Conn)
}

func NewClient(addr, port string) (*Client, error) {

	var (
		err error
		tcpAddr *net.TCPAddr
		conn *net.TCPConn
	)

	c := Client{}

	c.quit = make(chan bool)

	c.Address = addr
	c.Port = port

	tcpAddr, err = net.ResolveTCPAddr("tcp", addr + ":" + port) //获取一个TCP地址信息,TCPAddr
	if err != nil {
		c.isAlive = false
		return nil, err
	}

	conn, err = net.DialTCP("tcp", nil, tcpAddr) //创建一个TCP连接:TCPConn
	if err != nil {
		c.isAlive = false
		return nil, err
	}

	c.isAlive = true
	c.Conn = conn
	c.maxRetry = 3

	return &c, nil
}


