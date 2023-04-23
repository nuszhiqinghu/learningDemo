package main

import (
	"context"
	"log"
	"net"

	pb "test_grpc/test_grpc/helloworld" // 导入生成的 helloworld.pb.go 文件

	"google.golang.org/grpc"
)

const (
	port = ":50051" // 端口号
)

type server struct{} // 定义 gRPC 服务的 server

// 实现 SayHello 方法，接受客户端的请求，返回 HelloResponse 响应
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	// 开启端口
	lis, err := net.Listen("tcp", port)
	// 如果监听失败，则退出程序并打印错误信息
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建一个 gRPC 服务实例
	s := grpc.NewServer()
	// 注册服务，把实现服务的 server 结构体绑定到 gRPC 服务实例中
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("start gRPC server on port %s\n", port)
	// .Server启动服务，监听lis
	// 如果监听失败，则退出程序并打印错误信息
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
