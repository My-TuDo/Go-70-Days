package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// 模拟将昨天的 logs_demo 文件夹打包要所，准备外发
	srcDir := "./logs_demo"
	targetFile := "backup_logs.tar.gz"

	fmt.Println("正在打包文件...")

	// 1.创建物理文件
	fw, err := os.Create(targetFile)
	if err != nil {
		fmt.Printf("创建压缩文件失败：%v\n", err)
		return
	}
	defer fw.Close()

	// 2.建立压缩管道（Gzip）
	// 作用：往 gw 里写数据，会自动压缩后流入 fw（硬盘）
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 3.建立打包管道（Tar）
	// 作用：往 tw 里写数据，它会自动记录文件名、权限信息；会自动打包后流入 gw（内存）
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// 4.线性连结：结合昨天的 Walk 逻辑，扫描并写入
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 生成在压缩包内部的路径（相对路径）
		header, _ := tar.FileInfoHeader(info, "")
		header.Name, _ = filepath.Rel(srcDir, path)

		// 5.写入文件头：告诉压缩包接下来是一个什么文件
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// 如果是目录，写完头就跳过
		if info.IsDir() {
			return nil
		}

		// 6.拷贝内容：真正的文件数据流转
		fr, _ := os.Open(path)
		defer fr.Close()

		// 像吸管一样，把 fr（源文件）的内容吸到 tw （压缩包）里
		io.Copy(tw, fr)

		fmt.Printf("已打包：%s\n", header.Name)
		return nil
	})

	fmt.Printf("任务完成！生成的备份包：%s\n", targetFile)
}
