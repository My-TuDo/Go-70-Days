package main

import (
	"context"
	"kvault/api/kvaultpb"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 简易客户端
	// 创建连接
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接服务器失败: %v", err)
	}
	defer conn.Close()

	// 创建 KVault 客户端
	client := kvaultpb.NewKVaultClient(conn)

	// 调 RPC
	ctx := context.Background()

	if len(os.Args) < 2 {
		log.Fatalf("用法: %s <set|get> [key] [value]", os.Args[0])
		return
	}
	switch os.Args[1] {
	case "set":
		setResp, err := client.Set(ctx, &kvaultpb.SetRequest{Key: os.Args[2], Value: os.Args[3]})
		if err != nil {
			log.Fatalf("调用 Set 失败: %v", err)
		}
		log.Printf("Set 响应: %v", setResp)
	case "get":
		getResp, err := client.Get(ctx, &kvaultpb.GetRequest{Key: os.Args[2]})
		if err != nil {
			log.Fatalf("调用 Get 失败: %v", err)
		}
		log.Printf("Get 响应: %v", getResp)
	case "delete":
		delResp, err := client.Delete(ctx, &kvaultpb.DeleteRequest{Key: os.Args[2]})
		if err != nil {
			log.Fatalf("调用 Delete 失败: %v", err)
		}
		log.Printf("Delete 响应: %v", delResp)
	case "list":
		listResp, err := client.ListKeys(ctx, &kvaultpb.ListKeysRequest{})
		if err != nil {
			log.Fatalf("调用 List 失败: %v", err)
		}
		log.Printf("List 响应: %v", listResp)
	default:
		log.Fatalf("未知命令: %s", os.Args[1])
	}
}
