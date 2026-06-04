package main

import (
	"kvault/api/kvaultpb"
	"kvault/internal/server"
	"kvault/internal/store"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := store.New()                              // 创建存储实例
	KvSrv := server.New(s)                        // 创建服务器实例，注入存储
	grpcSrv := grpc.NewServer()                   // 创建 gRPC 服务器
	kvaultpb.RegisterKVaultServer(grpcSrv, KvSrv) // 注册 KVault 服务 到 gRPC 服务器

	lis, err := net.Listen("tcp", ":50051") // 监听 TCP 端口
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}
	grpcSrv.Serve(lis) // 启动 gRPC 服务器
}
