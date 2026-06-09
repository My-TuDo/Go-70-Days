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
