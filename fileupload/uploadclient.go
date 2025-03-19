package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"sync"
	"time"
)

// 前端的实现流程和这个类似

const chunkSize = 100 * 1024 // 100k

func uploadFile(filePath, uploadURL string) error {
	// 进度信息
	var totalUploadBytes int64 = 0
	uploadBytesCh := make(chan int64, 1)
	done := make(chan struct{})

	// 查询文件上传情况
	filename := path.Base(filePath)
	// 拼接url
	url := fmt.Sprintf("%s?%s=%s", uploadURL, "filename", filename)
	// 发起请求，查询文件上传信息
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	mp := map[string]interface{}{}
	body, err := io.ReadAll(resp.Body)
	json.Unmarshal(body, &mp)
	// 解析数据，得到已上传的字节数，即获取到上传进度，发送到通道，进度条是根据通道信息进行展示
	uploadBytesCh <- int64(mp["upload_bytes"].(float64))
	// 获取已上传分片 数组
	uploadChunk := mp["upload_chunk"].([]interface{})

	// open file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// get file size
	fileIno, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileIno.Size()
	chunkNum := fileSize / chunkSize
	if fileSize%chunkSize != 0 {
		chunkNum += 1
	}

	// 如果上传分片数组 为空，去初始化
	if len(uploadChunk) == 0 {
		uploadChunk = make([]interface{}, chunkNum)
	}

	// 启动协程做进度监控
	go func() {
		var lastPercent float64 = -1
		for {
			select {
			case uploadBytes := <-uploadBytesCh:
				if uploadBytes >= totalUploadBytes {
					totalUploadBytes = uploadBytes
					currentPercent := float64(uploadBytes) / float64(fileSize) * 100
					if currentPercent != lastPercent {
						fmt.Printf("上传进度: %d 字节，%.2f%%\n", uploadBytes, currentPercent)
						lastPercent = currentPercent
					}
				}
			case <-done:
				return
			}
		}
	}()
	wg := sync.WaitGroup{}
	var seq int64
	// 遍历分片数，计算偏移量  08:24
	for seq = 0; seq < chunkNum; seq++ {
		/* 验证断点续传的测试方法
		if seq == 1 {
			continue
		}

		// 13579 跳过，上传的就是246810
		if seq == 3 || seq == 5 || seq == 7 || seq==9 {
			continue
		}
		*/

		var v int64
		if uploadChunk[seq] != nil {
			v = int64(uploadChunk[seq].(float64))
		}

		if v == 0 {
			offset := seq * chunkSize
			if _, err := file.Seek(offset, io.SeekStart); err != nil {
				return err
			}

			bytes1 := make([]byte, chunkSize)
			n, err := file.Read(bytes1)
			if err != nil {
				log.Fatal(err)
			}
			wg.Add(1)

			// 启动协程，向server端发送数据
			go func(seq int64, data []byte) {
				defer wg.Done()
				payload := bytes.NewReader(data)
				client := &http.Client{}
				req, err := http.NewRequest("POST", url, payload)
				if err != nil {
					fmt.Println(err)
					return
				}
				req.Header.Set("X-File-Name", filename)
				req.Header.Set("Content-Range", fmt.Sprintf("%d/%d/%d", seq, chunkNum, fileSize))
				req.Header.Set("Content-Type", "application/json")

				res, err := client.Do(req)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer res.Body.Close()

				body, err := io.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				mp := make(map[string]float64)
				json.Unmarshal(body, &mp)
				uploadBytesCh <- int64(mp["upload_bytes"]) // 发送完成后得到服务端的返回结果

			}(seq, bytes1[:n])
		}
	}
	wg.Wait()
	// 添加一个额外的同步机制，确保最后的进度更新已经处理完成
	// 发送最后一次进度更新，确保显示100%
	uploadBytesCh <- fileSize
	// 等待一小段时间，确保进度协程处理完最后的消息
	time.Sleep(100 * time.Millisecond)

	// 先关闭done通道，让进度监控协程退出
	close(done)
	// 再关闭uploadBytesCh通道
	close(uploadBytesCh)

	fmt.Println("file upload success...")
	return nil
}

// 5:49 https://www.bilibili.com/video/BV11S91YyEeA
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	//filePath := "res/test.jpg"
	//filePath := "res/test.pptx"
	filePath := "res/test.pdf"
	uploadURL := "http://localhost:8080/upload"
	err := uploadFile(filePath, uploadURL)
	if err != nil {
		fmt.Println("upload error: ", err)
	} else {
		fmt.Println("upload success")
	}

	<-ctx.Done()

}
