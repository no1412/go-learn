package main

import (
	"fmt"
	"go-learn/ch22/server"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	// http
	//client, err := rpc.DialHTTP("tcp", "localhost:1234")
	// rpc
	//client, err := rpc.Dial("tcp", "localhost:1234")
	// jsonrpc
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := server.Args{A: 7, B: 8}
	var reply int
	err = client.Call("MathService.Add", args, &reply)
	if err != nil {
		log.Fatal("MathService.Add error:", err)
	}
	fmt.Printf("MathService.Add: %d+%d=%d", args.A, args.B, reply)
}
