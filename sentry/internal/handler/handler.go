package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"sentry/internal/aggregator"
)

// AlertHandler 处理 POST /api/alert
//
// 你需要实现：
//  1. 只接受 POST 方法，其他方法返回 405
//  2. 从请求体读取 JSON，解码成 aggregator.Alert
//  3. 调用 agg.ShouldSend(alert) 判断是否推送
//  4. 如果需要推送，用 slog.Info 输出报警信息
//  5. 返回 JSON 响应 { "pushed": true/false, "reason": "..." }
//
// 提示：
//   - json.NewDecoder(r.Body).Decode(&alert)
//   - 响应用 w.WriteHeader() + json.NewEncoder(w).Encode()
func AlertHandler(agg *aggregator.Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w 是 响应（写回给客户端的）
		// r 是 请求（客户端发过来的）
		// TODO: 实现报警处理逻辑
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed) // 405 方法不允许
			json.NewEncoder(w).Encode(map[string]any{
				"error":           "只支持 POST 方法",
				"reason":          "method_not_allowed",
				"allowed_methods": []string{http.MethodPost},
				"received_method": r.Method,
			})
			return // 405 方法不允许
		}

		// 解析请求体中的 JSON
		var alert aggregator.Alert // 定义一个 alert 变量来存储解码后的数据
		err := json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // 400 错误请求
			json.NewEncoder(w).Encode(map[string]any{
				"error":   "请求体必须是合法的 JSON",
				"reason":  "invalid_json",
				"details": err.Error(),
			})
			return // 400 错误请求
		}

		// 判断是否应该推送报警
		pushed := agg.ShouldSend(alert)
		if pushed {
			slog.Info("报警触发",
				"title", alert.Title,
				"level", alert.Level,
				"source", alert.Source,
			)
		}

		w.Header().Set("Content-Type", "application/json") // 设置响应头，告诉客户端这是 JSON 数据
		json.NewEncoder(w).Encode(map[string]any{
			"pushed": pushed,
			"reason": func() string {
				if pushed {
					return "new_alert"
				}
				return "duplicate_within_window"
			}(),
		})
	}
}

// json.NewDecoder(r.Body).Decode(&alert)
// NewDecoder 	// 新 decoder == “造一个翻译机”
// r.Body 		// 请求的身体 == "客户端发来的那堆 JSON 字符串"
// Decode 		// 解码 == "把 JSON 字符串翻译成 Go 结构体"
// &alert 		// 翻译结果放哪 == "存到 alert 这个变量里面"

// json.NewEncoder(w).Encode(...)
// NewEncoder 	// 新 encoder == “造一个翻译机（反向）”
// w 			// 响应的身体 == "要写回给客户端的那个地方"
// Encode 		// 编码 == "把 Go 结构体翻译成 JSON 字符串"
// ...			// 翻译内容 == "你想写回给客户端的数据"

// 反过来，编码 == "把 Go 结构体翻译成 JSON 字符串"
