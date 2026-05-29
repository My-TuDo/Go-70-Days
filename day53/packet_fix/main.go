package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// 1.[封包函数]：将消息包装成[长度+内容]
func pack(message string) []byte {
	// 获取消息长度
	length := int32(len(message))
	pkg := new(bytes.Buffer)

	// 写入长度 （使用 BigEndian 大端序）
	binary.Write(pkg, binary.BigEndian, length)
	// 写入内容
	binary.Write(pkg, binary.BigEndian, []byte(message))

	return pkg.Bytes()
}

// 2.[解包函数]：从[长度+内容]中提取消息
func unpack(reader *bufio.Reader) (string, error) {
	// A. 先读取长度（4字节）
	lengthByte, err := reader.Peek(4)
	if err != nil {
		return "", err
	}

	// B. 将长度字节转换为整数
	var length int32
	buffer := bytes.NewReader(lengthByte)
	binary.Read(buffer, binary.BigEndian, &length)

	// C. 检查缓冲区：如果还没攒够这么多数据，说明包还没传完，等待再试
	if int32(reader.Buffered()) < length+4 {
		return "", fmt.Errorf("数据未就绪")
	}

	// D. 读取完整包（长度 + 内容）
	pack := make([]byte, int(length+4))
	_, err = io.ReadFull(reader, pack)
	if err != nil {
		return "", err
	}

	return string(pack[4:]), nil
}

func main() {
	// 模拟服务端接收逻辑
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	fmt.Println("服务器已启动，监听9001端口...")

	go func() {
		for {
			conn, _ := listener.Accept()
			go handleConn(conn)
		}
	}()

	// --- 模拟客户端发送“粘包”数据 ---
	conn, _ := net.Dial("tcp", "127.0.0.1:9001")

	// 连续发送两个包， 中间不休息，强制造成粘包现象
	conn.Write(pack("Skill:Ultimate"))
	conn.Write(pack("Move:Forward"))
	conn.Close()

	select {} // 保持程序运行
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := unpack(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		fmt.Printf("[解析成功]收到独立指令：%s\n", msg)
	}
}
