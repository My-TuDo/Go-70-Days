package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"sentry/internal/aggregator"
	"sentry/internal/handler"
)

func main() {
	// ===== 1. 读取配置 =====
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	windowSec := os.Getenv("DEDUP_WINDOW")
	window := 5 * time.Minute
	if windowSec != "" {
		if sec, err := strconv.Atoi(windowSec); err == nil {
			window = time.Duration(sec) * time.Second
		}
	}

	agg := aggregator.New(window)

	// ===== 2. 注册路由 =====
	mux := http.NewServeMux()
	mux.HandleFunc("/api/alert", handler.AlertHandler(agg))

	// ===== 3. 启动 HTTP 服务 =====
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		slog.Info("服务启动", "port", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("服务启动失败", "error", err)
			os.Exit(1)
		}
	}()

	// ===== 4. 优雅退出 =====
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("正在关闭服务...")
	srv.Close()
	fmt.Println("服务已退出")
}
