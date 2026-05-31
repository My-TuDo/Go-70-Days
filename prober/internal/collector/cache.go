package collector

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"prober/internal/model"

	"github.com/redis/go-redis/v9"
)

// 提示：你需要用到的知识点
//
// redis.Client.Set — 写入 KV（参考文档）
// redis.Client.Get — 读取 KV
// json.Marshal / json.Unmarshal — 序列化/反序列化

// UpdateCache 将探测结果写入 Redis 缓存
//
// Redis key 格式: "probe:status:<目标名称>"
// 例如: probe:status:baidu
// 过期时间: 60 秒（TTL）
//
// 你需要实现：
//  1. 将 ProbeResult 转换成 ProbeCache（补充 CheckedAt 字段）
//  2. 用 json.Marshal 序列化
//  3. 用 rdb.Set 写入 Redis，设置 60 秒过期
//
// 提示：
//
//	ctx := context.Background()
//	jsonData, _ := json.Marshal(cacheData)
//	rdb.Set(ctx, "probe:status:baidu", jsonData, 60*time.Second)
func UpdateCache(rdb *redis.Client, result model.ProbeResult) {
	// TODO: 在这里写你的代码
	//

	cache := model.ProbeCache{
		TargetName: result.TargetName,
		TargetURL:  result.TargetURL,
		Success:    result.Success,
		StatusCode: result.StatusCode,
		LatencyMs:  result.LatencyMs,
		CheckedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	// 将 cache 转换成 JSON
	jsonData, _ := json.Marshal(cache)

	// 写入 Redis，设置 60 秒过期
	key := fmt.Sprintf("probe:status:%s", result.TargetName)
	ctx := context.Background()
	rdb.Set(ctx, key, jsonData, 60*time.Second)

}

// GetCachedStatus 从 Redis 读取某个目标的最新状态
//
// 你需要实现：
//  1. 用 rdb.Get 读取 key
//  2. 如果 key 不存在（redis.Nil），返回 nil
//  3. 用 json.Unmarshal 反序列化成 ProbeCache
//  4. 返回 &cache
func GetCachedStatus(rdb *redis.Client, targetName string) (*model.ProbeCache, error) {
	// TODO: 在这里写你的代码
	//
	key := fmt.Sprintf("probe:status:%s", targetName)
	ctx := context.Background()
	jsonData, err := rdb.Get(ctx, key).Bytes()
	if err == redis.Nil {
		fmt.Printf("缓存不存在: %s\n", key)
		return nil, nil
	}
	if err != nil {
		fmt.Printf("读取缓存失败: %v\n", err)
		return nil, err
	}

	var cache model.ProbeCache
	json.Unmarshal(jsonData, &cache)
	return &cache, nil

}
