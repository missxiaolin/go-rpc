package main

import (
	"net"
	"fmt"
	"go-rpc/services"
	"go-rpc/proto"
	"go-rpc/utils"
	"reflect"
)

func handle(conn *net.TCPConn) {

	var remoteAddr = conn.RemoteAddr() //获取连接到的对像的IP地址。
	fmt.Println("收到连接请求：", remoteAddr)
	fmt.Println("正在读取消息...")

	bys, _ := utils.ReadBytes(conn)

	fmt.Println("接收到客户端的消息：", string(bys))

	// 解析JSON数据
	data := proto.RequestBytes(bys)
	var service = data.Service
	var method = utils.StrFirstToUpper(data.Method)
	var arguments = data.Arguments

	//创建带调用方法时需要传入的参数列表
	parms := []reflect.Value{}
	for _, v := range arguments {
		parms = append(parms, reflect.ValueOf(v))
	}

	//使用方法名字符串调用指定方法
	var function = services.ServiceMappers[service][method]
	var result = function.Call(parms)
	fmt.Println("数据结果", result)

	conn.Write(proto.ResponseSuccess(result[0]))
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
