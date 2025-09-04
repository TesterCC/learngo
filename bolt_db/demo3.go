package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	// 打开 trivy.db
	db, err := bolt.Open("D:\\ws_go\\learngo\\bolt_db\\fanal\\trivy.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// print bucket
	err = db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			fmt.Println("bucket:", string(name))
			return nil
		})
	})
	if err != nil {
		log.Fatal(err)
	}

	// Trivy 用 "vulnerability" bucket 存漏洞详情
	err = db.View(func(tx *bolt.Tx) error {
		//b := tx.Bucket([]byte("artifact"))
		b := tx.Bucket([]byte("blob"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}

		// 例子：读取指定 CVE-2017-11329 信息
		vulnID := "CVE-2017-11329"
		val := b.Get([]byte(vulnID))
		if val != nil {
			fmt.Println(string(val)) // 这里是 JSON 格式
		} else {
			fmt.Println("not found:", vulnID)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
