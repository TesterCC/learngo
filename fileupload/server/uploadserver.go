package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

const uploadDir = "./uploads" // 上传文件存储目录
const tmp = "./tmp"

// 并发访问，多分片多线程上传
type File struct {
	lock        sync.RWMutex // 读写锁
	filename    string
	savePath    string
	tmpdir      string
	totalBytes  int64           // 总字节数
	uploadBytes int64           // 已上传字节数
	chunkNums   int64           // 分片数量
	uploadChunk map[int64]int64 // 已上传的分片信息
}

// 根据传入的序号和分片内容保存分片信息
func (f *File) UploadChunk(seq int64, reader io.Reader) (int64, int64) {
	err := os.MkdirAll(f.tmpdir, os.ModePerm) // 0777

	if err != nil {
		log.Fatal(err)
	}

	// 将内容写入临时文件
	file, err := os.OpenFile(fmt.Sprintf("%s/%d", f.tmpdir, seq), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	n, err := io.Copy(file, reader)

	if err != nil {
		file.Close()
		log.Fatal(err)
	}

	file.Close()

	// 加写文件锁
	f.lock.Lock()
	defer f.lock.Unlock()

	f.uploadChunk[seq] = n
	f.uploadBytes = 0

	for _, v := range f.uploadChunk {
		f.uploadBytes += v
	}
	if f.uploadBytes == f.totalBytes {
		err = f.saveFile()
		if err != nil {
			log.Fatal(err)
		}
		os.RemoveAll(f.tmpdir)
	}
	return int64(len(f.uploadChunk)), f.uploadBytes

}

func (f *File) saveFile() error {
	// ai gen
	// 确保上传目录存在
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return err
	}

	// 创建最终文件
	destFile, err := os.Create(f.savePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 按顺序合并所有分片
	for i := int64(0); i < f.chunkNums; i++ {
		chunkPath := fmt.Sprintf("%s/%d", f.tmpdir, i)

		// 打开分片文件
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			return err
		}

		// 将分片内容复制到目标文件
		_, err = io.Copy(destFile, chunkFile)
		chunkFile.Close()

		if err != nil {
			return err
		}
	}

	return nil

}

// 多个文件，多个分片并发
type FileInfoCache struct {
	lock sync.RWMutex
	data map[string]*File // 以文件名称为key缓存文件信息
}

var Cache = FileInfoCache{
	lock: sync.RWMutex{},
	data: map[string]*File{},
}

// 确保在使用缓存前，缓存里有文件信息
func (c *FileInfoCache) Init(filename, savePath, tmpdir string, chunkNums, totalBytes int64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	f, ok := c.data[filename]
	// 如果缓存中没有文件信息，就先做初始化
	if !ok {
		f = &File{
			filename:    filename,
			savePath:    savePath,
			tmpdir:      tmpdir,
			chunkNums:   chunkNums,
			totalBytes:  totalBytes,
			uploadChunk: map[int64]int64{},
		}
		c.data[filename] = f
	}

}

func (c *FileInfoCache) Get(filename string) *File {
	// 用的读锁，可以并发
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.data[filename]
}

// 处理文件上传
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleUpload(w, r)
	case http.MethodGet:
		handleResume(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// 处理文件上传请求
func handleUpload(w http.ResponseWriter, r *http.Request) {
	// 解析文件名和文件大小
	fileName := r.Header.Get("X-File-Name")
	if fileName == "" {
		http.Error(w, "File name is empty", http.StatusBadRequest)
		return
	}
	savePath := fmt.Sprintf("%s/%s", uploadDir, fileName)

	contentRange := r.Header.Get("Content-Range")
	if contentRange == "" {
		http.Error(w, "Content Range is empty", http.StatusBadRequest)
		return
	}

	// 解析 Content-Range 头部，格式为 "seq/nums/total"  分片信息
	rangeParts := strings.Split(contentRange, "/")
	if len(rangeParts) != 3 {
		log.Fatal("Content-Range Structure Error")
	}

	seq, err := strconv.ParseInt(rangeParts[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	nums, err := strconv.ParseInt(rangeParts[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	total, err := strconv.ParseInt(rangeParts[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	tmpDir := tmp + "/" + fileName

	Cache.Init(fileName, savePath, tmpDir, nums, total)

	f := Cache.Get(fileName)
	// 上传分片数据
	uploadChunkNum, uploadBytes := f.UploadChunk(seq, r.Body)
	defer r.Body.Close()

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	mp := map[string]int64{
		"upload_chunk_num": uploadChunkNum,
		"upload_bytes":     uploadBytes,
	}
	bytes, _ := json.Marshal(mp)
	w.Write(bytes)
}

// 处理断点续传请求
func handleResume(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("filename")
	if fileName == "" {
		http.Error(w, "File name is empty", http.StatusBadRequest)
		return
	}

	f := Cache.Get(fileName)

	mp := map[string]interface{}{
		"upload_bytes": 0,
		"upload_chunk": []int64{},
	}

	if f != nil {
		arr := make([]int64, f.chunkNums)

		mp["upload_bytes"] = f.uploadBytes
		for k, v := range f.uploadChunk {
			arr[k] = v
		}
		mp["upload_chunk"] = arr
	}

	bytes, _ := json.Marshal(mp)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func main() {
	// 注册文件上传处理函数
	http.HandleFunc("/upload", uploadHandler)
	// 启动HTTP服务器
	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
