package aggregator

import (
	"testing"
	"time"
)

// 新报警第一次来
func TestShouldSend_NewAlert(t *testing.T) {
	agg := New(5 * time.Minute) // 五分钟窗口

	alert := Alert{
		Title:  "服务宕机",
		Source: "prober",
	}

	result := agg.ShouldSend(alert)
	if !result {
		t.Errorf("新报警应返回 true，但返回了 false")
	}
}
