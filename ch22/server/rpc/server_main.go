package main

import (
	"go-learn/ch22/server"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("MathService", new(server.MathService))
	// 处理http
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// #1 tcp 接受请求
	// rpc.Accept(l)

	// #2 换成http的服务
	// http.Serve(l, nil)s

	// #3 json rpc
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("jsonrpc.Serve: accept:", err.Error())
			return
		}
		go jsonrpc.ServeConn(conn)
	}
}
