package main

import (
	"bufio" // 用于带缓冲的读取
	"fmt"
	"io" // 包含 io.ReadAll 等
	"os" // 包含文件操作、路径操作等
	"strings"
	"time"
)

func main() {
	fmt.Println("--- 第4周学习：常用标准库 (os 和 io 包 - 文件操作) ---")

	fileName := "example.txt"
	tempDir := "temp_dir_for_os_example"

	// --- 1. 文件写入 ---
	fmt.Println("\n--- 1. 文件写入 ---")

	// a) os.WriteFile (简单写入字节切片到文件，会创建或覆盖文件)
	//    权限 0666 表示所有用户可读写 (在Unix系统中)
	contentBytes := []byte("Hello from os.WriteFile!\nThis is a new line.\n")
	err := os.WriteFile(fileName, contentBytes, 0666)
	if err != nil {
		fmt.Printf("  os.WriteFile 错误: %v\n", err)
	} else {
		fmt.Printf("  成功使用 os.WriteFile 写入到 %s\n", fileName)
	}

	// b) os.Create / file.WriteString (更灵活的写入)
	file, err := os.Create("another_example.txt") // 创建文件，如果已存在则清空
	if err != nil {
		fmt.Printf("  os.Create 错误: %v\n", err)
	} else {
		defer file.Close() // 确保文件关闭
		bytesWritten, err := file.WriteString("Hello from file.WriteString!\n")
		if err != nil {
			fmt.Printf("  file.WriteString 错误: %v\n", err)
		} else {
			fmt.Printf("  成功使用 file.WriteString 写入 %d 字节到 %s\n", bytesWritten, file.Name())
		}
		// 也可以使用 file.Write([]byte(...))
	}

	// --- 2. 文件读取 ---
	fmt.Println("\n--- 2. 文件读取 ---")

	// a) os.ReadFile (简单读取整个文件内容到字节切片)
	readBytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("  os.ReadFile (%s) 错误: %v\n", fileName, err)
	} else {
		fmt.Printf("  os.ReadFile (%s) 读取内容:\n%s\n", fileName, string(readBytes))
	}

	// b) os.Open / io.ReadAll (更通用的读取)
	fileToRead, err := os.Open(fileName) // 只读方式打开文件
	if err != nil {
		fmt.Printf("  os.Open (%s) 错误: %v\n", fileName, err)
	} else {
		defer fileToRead.Close()
		allData, err := io.ReadAll(fileToRead) // 从 io.Reader 读取所有数据
		if err != nil {
			fmt.Printf("  io.ReadAll (%s) 错误: %v\n", fileName, err)
		} else {
			fmt.Printf("  io.ReadAll (%s) 读取内容:\n%s\n", fileName, string(allData))
		}
	}

	// c) os.Open / bufio.NewScanner (逐行读取，适合大文件)
	fileForScanner, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("  os.Open (%s) for scanner 错误: %v\n", fileName, err)
	} else {
		defer fileForScanner.Close()
		scanner := bufio.NewScanner(fileForScanner) // 创建扫描器
		fmt.Printf("  使用 bufio.Scanner 逐行读取 %s:\n", fileName)
		lineNum := 1
		for scanner.Scan() { // 每次调用 Scan 会读取一行，移除行尾的换行符
			fmt.Printf("    行 %d: %s\n", lineNum, scanner.Text()) // Text() 返回当前行的内容
			lineNum++
		}
		if err := scanner.Err(); err != nil { // 检查扫描过程中是否发生错误
			fmt.Printf("  扫描器错误: %v\n", err)
		}
	}

	// --- 3. 获取文件信息 (os.Stat) ---
	fmt.Println("\n--- 3. 获取文件信息 ---")
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("  os.Stat (%s) 错误: %v\n", fileName, err)
	} else {
		fmt.Printf("  文件信息 (%s):\n", fileName)
		fmt.Printf("    名称: %s\n", fileInfo.Name())
		fmt.Printf("    大小 (字节): %d\n", fileInfo.Size())
		fmt.Printf("    权限: %s\n", fileInfo.Mode().String()) // e.g., -rw-r--r--
		fmt.Printf("    修改时间: %s\n", fileInfo.ModTime().Format(time.RFC1123))
		fmt.Printf("    是否是目录: %t\n", fileInfo.IsDir())
	}

	// --- 4. 目录操作 (简单示例) ---
	fmt.Println("\n--- 4. 目录操作 ---")
	// a) 创建目录 (os.Mkdir, os.MkdirAll)
	//    os.Mkdir 创建单级目录，如果父目录不存在会失败
	//    os.MkdirAll 创建多级目录，类似 mkdir -p
	err = os.MkdirAll(tempDir, 0755) // 0755 权限 (rwxr-xr-x)
	if err != nil {
		fmt.Printf("  os.MkdirAll (%s) 错误: %v\n", tempDir, err)
	} else {
		fmt.Printf("  成功创建目录: %s\n", tempDir)
	}

	// b) 读取目录内容 (os.ReadDir)
	//    返回一个 []fs.DirEntry 切片 (fs.DirEntry 是 Go 1.16 引入的)
	dirEntries, err := os.ReadDir(".") // 读取当前目录 "."
	if err != nil {
		fmt.Printf("  os.ReadDir (\".\") 错误: %v\n", err)
	} else {
		fmt.Println("  当前目录 (\".\") 内容 (部分):")
		count := 0
		for _, entry := range dirEntries {
			if count < 5 || strings.HasPrefix(entry.Name(), "week") { // 只显示部分或特定文件
				entryType := "文件"
				if entry.IsDir() {
					entryType = "目录"
				}
				fmt.Printf("    %s: %s\n", entryType, entry.Name())
				count++
			}
		}
	}

	// --- 5. 删除文件和目录 (os.Remove, os.RemoveAll) ---
	fmt.Println("\n--- 5. 删除文件和目录 ---")
	// 删除文件
	// err = os.Remove("another_example.txt")
	// if err != nil {
	// 	// 在某些系统上（尤其是Windows），即使文件已关闭，立即删除也可能因“文件被占用”而失败。
	// 	// 为确保示例流程顺畅，此处注释掉。
	// 	fmt.Printf("  os.Remove (\"another_example.txt\") 错误: %v (此错误在某些系统上可能发生)\n", err)
	// } else {
	// 	fmt.Println("  成功删除文件: another_example.txt")
	// }

	// 删除目录 (os.Remove 只能删除空目录, os.RemoveAll 可以删除非空目录)
	err = os.RemoveAll(tempDir)
	if err != nil {
		fmt.Printf("  os.RemoveAll (%s) 错误: %v\n", tempDir, err)
	} else {
		fmt.Printf("  成功删除目录: %s\n", tempDir)
	}
	// 清理主示例文件
	// os.Remove(fileName) // 可以在测试后删除

	fmt.Println("\n--- os 和 io 包文件操作学习结束 ---")
}
