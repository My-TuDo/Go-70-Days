package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// 部署任务的结构体
type DeployTask struct {
	Host       string
	LocalPath  string
	RemotePath string
}

func main() {
	// 模拟生产环境的发布：将本地的 tar.gz 包推送到另一个服务器上
	task := DeployTask{
		Host:       "127.0.0.1:22",
		LocalPath:  "./backup_logs.tar.gz",
		RemotePath: "/tmp/backup_logs_deployed.gz",
	}

	// 1.连接远程服务器（SSH）
	sshConfig := &ssh.ClientConfig{
		User: "xzh",
		Auth: []ssh.AuthMethod{
			ssh.Password("1234567"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	fmt.Printf("正在建立 SSH 隧道：%s\n", task.Host)

	// 2.建立连接
	conn, err := ssh.Dial("tcp", task.Host, sshConfig)
	if err != nil {
		fmt.Printf("连接 SSH 失败：%v\n", err)
		return
	}
	defer conn.Close()

	// 3.创建 SFTP 客户端
	// 作用：通过 SSH 隧道进行文件传输
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Printf("建立 SFTP 会话失败：%v\n", err)
		return
	}
	defer sftpClient.Close()

	// 4.打开本地文件流
	srcFile, err := os.Open(task.LocalPath)
	if err != nil {
		fmt.Printf("读取本地文件失败：%v\n", err)
		return
	}
	defer srcFile.Close()

	// 5.创建远程文件句柄
	// 作用：在目标服务器上”占位“，准备写入
	dstFile, err := sftpClient.Create(task.RemotePath)
	if err != nil {
		fmt.Printf("创建远程文件失败：%v\n", err)
		return
	}
	defer dstFile.Close()

	// 6.文件传输,使用流式拷贝（Zero-Copy）
	fmt.Println("正在上传...")
	start := time.Now()
	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Printf("文件传输失败：%v\n", err)
		return
	}

	fmt.Printf("发布成功！\n传输大小：%.2f MB | 耗时：%v\n", float64(bytesCopied)/1024/1024, time.Since(start))
}
