package main

import (
	"net"
	"fmt"
	"services"
	"proto"
	"utils"
)

func handle(conn *net.TCPConn) {

	var remoteAddr = conn.RemoteAddr() //获取连接到的对像的IP地址。
	fmt.Println("收到连接请求：", remoteAddr)
	fmt.Println("正在读取消息...")

	bys, _ := utils.ReadBytes(conn)

	fmt.Println("接收到客户端的消息：", string(bys))


}

func main()  {
	localAddress, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:11520") //定义一个本机IP和端口。
	var tcpListener, err = net.ListenTCP("tcp", localAddress)       //在刚定义好的地址上进监听请求。
	if err != nil {
		fmt.Println("监听出错：", err)
		return
	}
	defer func() { //担心return之前忘记关闭连接，因此在defer中先约定好关它。
		tcpListener.Close()
	}()
	services.Init()
	fmt.Println("正在等待连接...")
	for {
		var conn, err2 = tcpListener.AcceptTCP() //接受连接。
		if err2 != nil {
			fmt.Println("接受连接失败：", err2)
			return
		}

		go handle(conn)
	}
}
