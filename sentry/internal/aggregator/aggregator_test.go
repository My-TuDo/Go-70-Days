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

// 同一报警在窗口内再次来
func TestShouldSend_Duplicate(t *testing.T) {
	agg := New(5 * time.Minute)

	alert := Alert{
		Title:  "服务宕机",
		Source: "prober",
	}

	agg.ShouldSend(alert) // 第一次推送

	result := agg.ShouldSend(alert) // 第二次推送
	if result {
		t.Errorf("同一报警在窗口内应返回 false，但返回了 true")
	}
}

// 不同报警
func TestShouldSend_DifferentAlert(t *testing.T) {
	agg := New(5 * time.Minute)

	alerts := []Alert{
		{
			Title:  "服务宕机",
			Source: "prober",
		},
		{
			Title:  " 磁盘告警",
			Source: "prometheus",
		},
	}

	ag.ShouldSend(alerts[0]) // 推送第一个报警

	result := agg.ShouldSend(alerts[1]) // 推送第二个报警
	if !result {
		t.Errorf("不同报警应返回 true，但返回了 false")
	}
}
