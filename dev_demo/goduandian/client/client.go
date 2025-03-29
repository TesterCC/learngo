package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"learngo/dev_demo/goduandian/util"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// 定义客户端文件存储路径
// 所有从服务器下载的文件都将保存在这个目录下
var (
	FilePath string
)

// Download 函数实现文件的断点续传下载功能
// 参数:
//   - FileName: 要下载的文件名，该文件应存在于服务器的文件目录中
func Download(FileName string) {
	// 打开或创建本地文件，准备写入下载的数据
	// os.O_CREATE: 如果文件不存在则创建
	// os.O_APPEND: 以追加模式打开，新数据将添加到文件末尾（这是断点续传的关键）
	// os.O_WRONLY: 以只写模式打开
	// os.ModePerm: 设置文件权限为0777（所有用户可读写执行）

	fmt.Printf("[D]Download file info: %s%s\n", FilePath, FileName)

	fout, _ := os.OpenFile(FilePath+FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	// 确保函数结束时关闭文件，防止资源泄露
	defer fout.Close()

	// 获取已下载文件的当前大小，这将作为断点位置
	stat, _ := fout.Stat()
	FileSize := stat.Size()

	// 构造请求参数，包含文件名和断点位置
	// StartPoint表示已下载的字节数，服务器将从这个位置开始传输
	request := util.GetRequest{
		FileName:   FileName,
		StartPoint: int(FileSize),
	}

	// 将请求参数序列化为JSON格式
	bs, _ := json.Marshal(request)

	// 创建HTTP POST请求
	// 使用127.0.0.1:1234作为服务器地址
	// 请求体是序列化后的JSON数据
	req, err := http.NewRequest("POST", "http://127.0.0.1:1234/download", bytes.NewBuffer(bs))
	if err != nil {
		// 如果创建请求失败，记录错误并退出函数
		log.Printf("create request err: %v", err)
		return
	}
	// 设置Content-Type头，告知服务器请求体是JSON格式
	req.Header.Set("Content-Type", "application/json")

	// 创建HTTP客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// 如果发送请求失败，记录错误并退出函数
		log.Printf("http post request err: %v", err)
		return
	}
	// 确保函数结束时关闭响应体，防止资源泄露
	defer resp.Body.Close()

	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		// 如果状态码不是200 OK，记录错误并将响应体内容输出到标准输出
		log.Printf("status error: %d", resp.StatusCode)
		io.Copy(os.Stdout, resp.Body)
		return
	}

	// 从响应头中获取文件总大小
	// 默认设置为1MB，防止服务器没有返回文件大小
	FileTotalSize := int64(1 << 20) // 1MB = 2^20 字节
	if v, exists := resp.Header["File-Size"]; exists {
		// 如果响应头中包含File-Size字段，解析为整数
		size, err := strconv.ParseInt(v[0], 10, 64)
		if err != nil {
			// 如果解析失败，记录错误并使用默认值
			log.Printf("parse file size err: %v", err)
			FileTotalSize = int64(1 << 20) // 默认1MB
		} else {
			// 解析成功，使用服务器返回的文件大小
			FileTotalSize = size
		}
	}

	// 初始化进度条
	// 参数1: 已下载的字节数（断点位置）
	// 参数2: 文件总大小
	bar := util.Bar{}
	bar.NewOption(FileSize, int64(FileTotalSize))

	// 记录本次下载的字节数（不包括之前已下载的部分）
	accDownload := int64(0)

	// 使用缓冲区读取响应体并写入本地文件
	buffer := make([]byte, 1024) // 创建1KB的缓冲区
	for {
		// 从响应体读取数据到缓冲区
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			// 将读取到的数据写入本地文件
			_, writeErr := fout.Write(buffer[:n])
			if writeErr != nil {
				// 如果写入失败，记录错误并跳出循环
				log.Printf("write file err: %v", writeErr)
				break
			}
			// 更新已下载字节数并刷新进度条
			accDownload += int64(n)
			// 参数是总下载量（本次下载量+之前已下载量）
			bar.Play(accDownload + FileSize)
		}
		if err != nil {
			if err != io.EOF {
				// 如果不是文件结束错误，记录异常
				log.Printf("read resp body exception: %s", err)
			} else {
				// 如果是文件结束（下载完成），打印换行符
				// 这是为了让进度条显示完整后，后续日志能正常显示在新行
				fmt.Println()
			}
			// 无论是正常结束还是异常，都跳出循环
			break
		}
	}
}

// main 函数是客户端程序的入口点
// 启动下载过程，尝试下载名为"test.mp4"的文件
func main() {
	// 初始化下载路径为当前工作目录下的localdownload文件夹
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前工作目录失败: %v", err)
	}
	FilePath = filepath.Join(cwd, "localdownload") + "/"
	fmt.Printf("文件将被下载到: %s\n", FilePath)

	// 确保下载目录存在
	err = os.MkdirAll(FilePath, os.ModePerm)
	if err != nil {
		log.Fatalf("创建下载目录失败: %v", err)
	}

	// 调用Download函数开始下载文件
	// 如果之前已经下载了部分文件，将从断点处继续下载
	Download("test.mp4")
}
