package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	// 模拟在目标服务器执行：检查文件-》创建目录-》解压包
	config := &ssh.ClientConfig{
		User: "xzh",
		Auth: []ssh.AuthMethod{
			ssh.Password("1234567"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		log.Fatalf("无法连接服务器：%v", err)
		return
	}
	defer client.Close()

	// 建立会话
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("无法创建会话：%v", err)
		return
	}
	defer session.Close()

	// 2.构造命令链
	// 逻辑：
	/*
	* a. [-f /tmp/backup_logs_deployed.tar.gz]：检查文件是否存在
	* b. mkdir -p /tmp/backup_logs：创建解压目录
	* c. tar -zxvf ... -C ...： 执行解压
	 */
	// 完美的修复方案：
	remoteCmd := " [ -f /tmp/backup_logs_deployed.gz ] && mkdir -p /tmp/extracted_logs && gunzip -c /tmp/backup_logs_deployed.gz > /tmp/extracted_logs/backup_logs_deployed"

	fmt.Println("正在下达远程解压指令...")

	// 3.执行命令并获取实时回执
	// 使用 CombinedOutput 抓取解压过程中的每一个文件名
	output, err := session.CombinedOutput(remoteCmd)
	if err != nil {
		fmt.Printf("部署失败，原因可能为文件缺失或权限不足：%v\n", err)
		fmt.Println("调试信息：", string(output))
		return
	}

	fmt.Println("远程解压完成！回执清单如下：")
	fmt.Println(string(output))

	// 4.后置验证
	verifyRemoteDeployment(client)
}

func verifyRemoteDeployment(client *ssh.Client) {
	session, _ := client.NewSession()
	defer session.Close()

	// 统计解压出了多少个文件
	out, _ := session.Output("ls /tmp/extracted_logs | wc -l")
	fmt.Println("解压出的文件数量：", string(out))
}
