package main

import (
	"encoding/json"
	"io"
	"learngo/dev_demo/goduandian/util"
	"log"
	"net/http"
	"os"
	"strconv"
)

// 定义服务端文件存储路径
// 所有需要被客户端下载的文件都应该放在这个目录下
var (
	FilePath = "./files/"
)

// Download 处理客户端的断点续传下载请求
// 该函数是断点续传的核心实现，从客户端指定的断点位置开始传输文件
// 参数:
//   - w: HTTP响应写入器，用于向客户端发送数据
//   - r: HTTP请求对象，包含客户端发送的请求信息
func Download(w http.ResponseWriter, r *http.Request) {
	// 读取完整请求体（这里假设请求体很小，实际生产环境可能需要限制大小）
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		// 如果读取请求体失败，返回500错误
		http.Error(w, "读取请求体失败", http.StatusInternalServerError)
		return
	}

	// 将JSON请求体反序列化为GetRequest结构体
	// GetRequest包含文件名和断点位置信息
	var request util.GetRequest
	err = json.Unmarshal(bs, &request)
	if err != nil {
		// 如果JSON格式不正确，返回400错误
		http.Error(w, "请求体不合法", http.StatusBadRequest)
		return
	}

	// 打开客户端请求的文件
	// 文件路径由FilePath和请求中的文件名组成
	fin, err := os.Open(FilePath + request.FileName)
	if err != nil {
		// 如果文件不存在，返回404错误
		http.Error(w, "想下载的文件不存在", http.StatusNotFound)
		return
	}

	// 确保函数结束时关闭文件，防止资源泄露
	defer fin.Close()

	// 获取文件信息，主要是为了获取文件大小
	stat, _ := fin.Stat()
	FileSize := stat.Size()
	// 通过HTTP响应头告知客户端文件的总大小
	// 客户端可以用这个信息来计算下载进度
	w.Header().Add("File-Size", strconv.FormatInt(FileSize, 10))

	// 使用Seek函数将文件指针移动到断点位置
	// 这是断点续传的关键：从上次中断的地方继续传输
	// 第二个参数0表示从文件开始位置计算偏移量
	fin.Seek(int64(request.StartPoint), 0)
	log.Printf("") // 这行似乎是多余的，可能是为了日志格式化

	// 使用缓冲区读取文件并发送给客户端
	// 这种方式可以有效处理大文件，避免一次性加载整个文件到内存
	buffer := make([]byte, 1024) // 创建1KB的缓冲区
	for {
		// 读取文件内容到缓冲区
		n, err := fin.Read(buffer)
		if n > 0 {
			// 将读取到的数据写入HTTP响应
			// 只写入实际读取到的字节数(n)
			_, err = w.Write(buffer[0:n])
		}
		if err != nil {
			// 如果读取过程中出现错误
			if err != io.EOF {
				// 如果不是文件结束错误，记录异常
				log.Printf("给client返回文件时读文件异常： %s", err)
			}
			// 无论是正常结束还是异常，都跳出循环
			break
		}
	}
}

// main 函数是服务端程序的入口点
// 负责初始化服务器环境并启动HTTP服务
func main() {
	// 确保文件存储目录存在，如果不存在则创建
	// os.ModePerm表示创建的目录权限为0777（所有用户可读写执行）
	err := os.MkdirAll(FilePath, os.ModePerm)
	if err != nil {
		// 如果目录创建失败，记录错误并终止程序
		log.Fatalf("创建目录失败: %v", err)
	}

	// 创建HTTP路由复用器(ServeMux)
	// 用于将不同的URL路径映射到对应的处理函数
	mux := http.NewServeMux()
	// 注册下载处理函数，当客户端访问/download路径时会调用Download函数
	mux.HandleFunc("/download", Download)

	// 启动HTTP服务器，监听1234端口
	log.Printf("服务器启动在 http://localhost:1234")
	// ListenAndServe会阻塞当前goroutine，直到服务器关闭
	if err := http.ListenAndServe(":1234", mux); err != nil {
		// 如果服务器启动失败，记录错误并终止程序
		log.Fatalf("服务器启动失败: %v", err)
	}
}
