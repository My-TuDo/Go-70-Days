package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config 映射配置的结构体
// 这里的 mapstructure 标签不仅对应 YAML，也对应 Viper 转换后的环境变量
type Config struct {
	AppPort string `mapstructure:"app_port"`
	DBUser  string `mapstructure:"db_user"`
}

func main() {
	// 初始化一个 Viper 实例
	v := viper.New()

	// 1. 设置环境变量的前缀（比较 MHY_APP_PORT)
	// 设置前缀的目的：防止多个程序共用一台服务器时环境变量重名（冲突）
	// 例如：MHY_APP_PORT 和 OTHER_APP_PORT 就不会冲突了
	v.SetEnvPrefix("MHY")

	// 自动读取所有的 MHY_ 开头的变量，及开启自动绑定
	v.AutomaticEnv()

	// 2. 设置默认值
	// 即使运维没有设置环境变量， 程序也要能跑起来
	v.SetDefault("app_port", "8080")
	v.SetDefault("db_user", "root")

	// 3. 模拟在 Liunx 终端设置环境变量
	// 实际上在 K8s 部署时， 这些环境变量是在 YAML 部署文件中定义的
	os.Setenv("MHY_APP_PORT", "9999")
	os.Setenv("MHY_DB_USER", "sre_admin")

	// 将读取到的配置（环境变量 > 默认值）解压到结构体中
	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		fmt.Printf("解压配置失败: %v\n", err)
		return
	}

	fmt.Println("SRE 配置中心启动...")
	fmt.Printf("监听端口：%s\n", conf.AppPort)
	fmt.Printf("数据库用户：%s\n", conf.DBUser)
}
