package jsonrpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func StartServer() {
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	arith := new(Arith)
	err = rpc.Register(arith)
	if err != nil {
		log.Fatal("rpc register err:", err)
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(accept))
	}

}

func Client() {
	dial, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dial error:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(dial))

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Mutiply error", err)
	}
	fmt.Printf("Mutiply: %d*%d = %d \n", args.A, args.B, reply)

}
