package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string `mapstructure:"app_port"`
	DBUser  string `mapstructure:"db_user"`
}

func main() {
	v := viper.New()

	// 1. 设置环境变量的前缀（比较 MHY_APP_PORT)
	v.SetEnvPrefix("MHY")
	v.AutomaticEnv() // 2. 自动读取所有的 MHY_ 开头的变量

	// 2. 设置默认值
	v.SetDefault("app_port", "8080")
	v.SetDefault("db_user", "root")

	// 3. 模拟在 Liunx 终端设置环境变量
	os.Setenv("MHY_APP_PORT", "9999")
	os.Setenv("MHY_DB_USER", "sre_admin")

	var conf Config
	v.Unmarshal(&conf)

	fmt.Println("SRE 配置中心启动...")
	fmt.Printf("监听端口：%s\n", conf.AppPort)
	fmt.Printf("数据库用户：%s\n", conf.DBUser)
}
