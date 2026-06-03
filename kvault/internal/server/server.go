package server

import (
	"context"
	"kvault/api/kvaultpb"
	"kvault/internal/store"
)

type KVaultServer struct {
	kvaultpb.UnimplementedKVaultServer              // 嵌入未实现的服务器接口，确保前向兼容
	store                              *store.Store // 内部使用的存储实例
}

// New 创建一个新的 KVaultServer 实例，初始化存储
func New(store *store.Store) *KVaultServer {
	return &KVaultServer{
		store: store,
	}
}

// Set 写入键值对
func (ks *KVaultServer) Set(ctx context.Context, req *kvaultpb.SetRequest) (*kvaultpb.SetResponse, error) {
	ks.store.Set(req.GetKey(), req.GetValue()) // 调用存储的 Set 方法存储数据
	return &kvaultpb.SetResponse{Success: true}, nil
}

// Get 读取键值对
func (ks *KVaultServer) Get(ctx context.Context, req *kvaultpb.GetRequest) (*kvaultpb.GetResponse, error) {
	value, ok := ks.store.Get(req.GetKey()) // 调用存储的 Get 方法获取数据
	if !ok {
		return &kvaultpb.GetResponse{Found: false}, nil // 如果键不存在，返回 nil
	}
	return &kvaultpb.GetResponse{Value: value, Found: true}, nil // 返回获取到的值
}

// Delete 删除键值对
func (ks *KVaultServer) Delete(ctx context.Context, req *kvaultpb.DeleteRequest) (*kvaultpb.DeleteResponse, error) {
	ks.store.Delete(req.GetKey()) // 调用存储的 Delete 方法删除数据
	return &kvaultpb.DeleteResponse{Success: true}, nil
}

// ListKeys 列出所有键
func (ks *KVaultServer) ListKeys(ctx context.Context, req *kvaultpb.ListKeysRequest) (*kvaultpb.ListKeysResponse, error) {
	keys := ks.store.ListKeys() // 调用存储的 ListKeys 方法获取所有键
	return &kvaultpb.ListKeysResponse{Keys: keys}, nil
}
