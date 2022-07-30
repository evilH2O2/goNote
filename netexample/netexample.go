package netexample

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var listenWg sync.WaitGroup

func conDailTcp() {
	defer listenWg.Done()
	i := net.ParseIP("127.0.0.1")
	fmt.Printf("i.String(): %v\n", i.String())

	time.Sleep(time.Second * 2)
	// 建立 tcp 连接
	// 返回 net.Con 句柄
	c, err := net.DialTimeout("tcp", "127.0.0.1:4444", time.Millisecond*100)
	if err != nil {
		fmt.Printf("connect err: %v\n", err)
	} else {
		s := c.RemoteAddr().String()
		fmt.Printf("connect remote, address: %v\n", s)
		// 分配数据缓冲区
		buf := make([]byte, 11)
		// 读取数据到缓冲区
		n, err2 := c.Read(buf)
		if err != nil {
			log.Println("read err", err2)
		} else {
			fmt.Printf("read size: %d, data: %s\n", n, string(buf))
		}
		c.Close()
	}
}

func conListentcp() {
	defer listenWg.Done()
	l, err := net.Listen("tcp", "127.0.0.1:4444")

	if err != nil {
		log.Fatalf("connect err: %v\n", err)
	} else {
		fmt.Println("Listen start...")
		// 等到连接
		c, err2 := l.Accept()
		if err != nil {
			log.Printf("accept err: %v\n", err2)
		}
		// 写数据
		_, err3 := c.Write([]byte("hello world"))
		if err != nil {
			log.Println("write connect err: ", err3)
		}
		c.Close()
		l.Close()
		fmt.Println("Listen done")
	}
}

func InvokeNetLib() {
	listenWg.Add(2)
	go conListentcp()
	conDailTcp()

	listenWg.Wait()
}
