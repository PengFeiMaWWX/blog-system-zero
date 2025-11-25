package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

// ConfigManager configManager 统一管理etcd的操作
type ConfigManager struct {
	client *clientv3.Client
}

// NewConfigManager 创建管理器并初始化 etcd 客户端连接
func NewConfigManager(endpoints []string) (*ConfigManager, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second, // 设置连接超时时间
	})
	if err != nil {
		return nil, err
	}

	return &ConfigManager{client: client}, nil
}

// GetDBConfig 获取并解析数据库配置
func (cm *ConfigManager) GetConfig(key string) (string, error) {
	// 创建带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保上下文被取消，释放资源

	resp, err := cm.client.Get(ctx, key)
	if err != nil {
		return "", err
	}

	if len(resp.Kvs) == 0 {
		return "", nil
	}
	return string(resp.Kvs[0].Value), nil
}

func (cm *ConfigManager) GetDbConfig(key, value string) (string, error) {

	configStr, err := cm.GetConfig(key)
	if err != nil {
		return "", err
	}
	log.Printf("从 ETCD 获取的原始字符串: %s", configStr)
	log.Printf("字符串长度: %d", len(configStr))
	log.Printf("前10个字符: %q", configStr[:min(10, len(configStr))])

	// 检查是否是有效的 JSON
	if !json.Valid([]byte(configStr)) {
		log.Printf(" 无效的 JSON 格式")
		// 打印每个字符的 ASCII 值来诊断问题
		for i, ch := range configStr {
			log.Printf("字符[%d]: %q (ASCII: %d)", i, ch, ch)
		}
		return "", fmt.Errorf("无效的 JSON 格式")
	}

	// 解析数据库配置
	var dbConfig struct {
		DSN string `json:"dsn"`
	}
	if err := json.Unmarshal([]byte(configStr), &dbConfig); err != nil {
		log.Printf(" JSON 解析失败: %v", err)
		return "", err
	}

	log.Printf(" 成功解析 DSN: %s", dbConfig.DSN)

	return dbConfig.DSN, nil
}
