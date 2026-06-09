package server

import (
	"context"
	"kvault/api/kvaultpb"
	"kvault/internal/store"
	"testing"
)

func TestServerSetAndGet(t *testing.T) {
	s := store.New()
	svr := New(s)

	// Set
	setResp, err := svr.Set(context.Background(), &kvaultpb.SetRequest{
		Key:   "testKey",
		Value: "testValue",
	})
	// 期望： err == nil, setResp.Success == true
	if err != nil {
		t.Errorf("Set()返回错误: %v", err)
	}
	if !setResp.GetSuccess() {
		t.Errorf("Set()返回 Success=false，期望 true")
	}

	// Get
	getResp, err := svr.Get(context.Background(), &kvaultpb.GetRequest{
		Key: "testKey",
	})
	// 期望： err == nil, getResp.Found == true, getResp.Value == "testValue"
	if err != nil {
		t.Errorf("Get()返回错误: %v", err)
	}
	if !getResp.GetFound() {
		t.Errorf("Get()返回 Found=false，期望 true")
	}
	if getResp.GetValue() != "testValue" {
		t.Errorf("Get()返回 Value=%s，期望 testValue", getResp.GetValue())
	}
}

// 测试 Get 不存在的 key
func TestServerGetMissing(t *testing.T) {
	s := store.New()
	svr := New(s)

	getResp, err := svr.Get(context.Background(), &kvaultpb.GetRequest{
		Key: "missingKey",
	})
	// 期望： err == nil, getResp.Found == false
	if err != nil {
		t.Errorf("Get()返回错误: %v", err)
	}
	if getResp.GetFound() {
		t.Errorf("Get()返回 Found=true，期望 false")
	}
}

// 测试 Delete
func TestServerDelete(t *testing.T) {
	s := store.New()
	svr := New(s)

	// 先 Set 一个 key
	_, err := svr.Set(context.Background(), &kvaultpb.SetRequest{
		Key:   "keyToDelete",
		Value: "valueToDelete",
	})
	if err != nil {
		t.Errorf("Set()返回错误: %v", err)
	}

	// Delete
	delResp, err := svr.Delete(context.Background(), &kvaultpb.DeleteRequest{
		Key: "keyToDelete",
	})
	// 期望： err == nil, delResp.Success == true
	if err != nil {
		t.Errorf("Delete()返回错误: %v", err)
	}
	if !delResp.GetSuccess() {
		t.Errorf("Delete()返回 Success=false，期望 true")
	}

	// 再 Get 一次，应该找不到
	getResp, err := svr.Get(context.Background(), &kvaultpb.GetRequest{
		Key: "keyToDelete",
	})
	if err != nil {
		t.Errorf("Get()返回错误: %v", err)
	}
	if getResp.GetFound() {
		t.Errorf("Get()返回 Found=true，期望 false")
	}
}

// 测试 ListKeys
func TestServerListKeys(t *testing.T) {
	s := store.New()
	svr := New(s)

	// Set 多个 key
	keys := []string{"key1", "key2", "key3"}
	for _, key := range keys {
		_, err := svr.Set(context.Background(), &kvaultpb.SetRequest{
			Key:   key,
			Value: "someValue",
		})
		if err != nil {
			t.Errorf("Set()返回错误: %v", err)
		}
	}

	// ListKeys
	listResp, err := svr.ListKeys(context.Background(), &kvaultpb.ListKeysRequest{})
	// 期望： err == nil, listResp.Keys 包含所有设置的 keys
	if err != nil {
		t.Errorf("ListKeys()返回错误: %v", err)
	}

	keySet := make(map[string]struct{})
	for _, k := range listResp.GetKeys() {
		keySet[k] = struct{}{}
	}

	for _, expectedKey := range keys {
		if _, exists := keySet[expectedKey]; !exists {
			t.Errorf("ListKeys()返回的 keys 中缺少 %q", expectedKey)
		}
	}
}
