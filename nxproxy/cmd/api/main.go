package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", helloHandler)
	mux.HandleFunc("/api/status", statusHandler)
	mux.HandleFunc("/api/panic", panicHandler) // 用于测试 Nginx 502

	srv := &http.Server{Addr: ":50051", Handler: mux}

	go func() {
		slog.Info("API 服务启动", "port", 50051)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			slog.Error("服务启动失败", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("正在关闭...")
	srv.Close()
}

// helloHandler 返回 {"message": "Hello, nxproxy!"}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, nxproxy!"})
}

// statusHandler 返回服务状态和当前时间
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
		"time":   time.Now().Format("2006-01-02 15:04:05"),
	})
}

// panicHandler 模拟崩溃，Nginx 应返回 502
func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("模拟崩溃")
}
