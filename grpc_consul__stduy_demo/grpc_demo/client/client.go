package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "test_grpc/test_grpc/helloworld" // 这里修改为实际的 proto 文件路径

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	//连接到server端，此处禁用安全连接，没有加密和验证
	//conn是连接对象
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//所有的连接都要进行关闭
	defer conn.Close()

	//建立连接 创立了一个具有连接的c客户端 建立连接之后就可以直接调用服务端的代码
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//执行rpc调用
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
