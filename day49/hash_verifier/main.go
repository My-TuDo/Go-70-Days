package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

// 数据指纹校验

// CalculateLocalHash 计算本地文件的 SHA256 哈希值
func CalculateLocalHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	// io.Copy 同样适用于计算 Hash，它会分块读取，不占内存
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
	// 验证 /tmp/backup_logs_deployed.gz 的完整性
	localFile := "./backup_logs.tar.gz"
	remotePath := "/tmp/backup_logs_deployed.gz"

	// 计算本地指纹
	localHash, err := CalculateLocalHash(localFile)
	if err != nil {
		log.Fatalf("计算本地文件哈希失败：%v\n", err)
	}
	fmt.Printf("本地文件指纹：%s\n", localHash)

	// 准备远程环境
	config := &ssh.ClientConfig{
		User: "xzh",
		Auth: []ssh.AuthMethod{
			ssh.Password("1234567"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, _ := ssh.Dial("tcp", "127.0.0.1:22", config)
	defer client.Close()

	// 下令远程计算指纹
	// 直接调用 Linux 原生的 sha256sum 命令
	session, _ := client.NewSession()
	defer session.Close()

	fmt.Println(" 正在请求远程服务器校验...")
	remoteCmd := fmt.Sprintf("sha256sum %s", remotePath)
	output, _ := session.Output(remoteCmd)

	// sha256sum 的输出格式通常是：“hash值 文件名”
	remoteHash := strings.Fields(string(output))[0]
	fmt.Printf("远程文件指纹：%s\n", remoteHash)

	// 双重对齐
	if localHash == remoteHash {
		fmt.Println("[一致]数据完整无损，可部署")
	} else {
		fmt.Println("[警告]指纹不匹配，文件可能损坏或被篡改，启动清理逻辑..")
		// 模拟：删除损坏的远程文件
		cleanupSession, _ := client.NewSession()
		cleanupSession.Run(fmt.Sprintf("rm -f %s", remotePath))
		cleanupSession.Close()
	}
}
