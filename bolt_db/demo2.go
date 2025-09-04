package main

import (
	"encoding/json"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	// 打开 trivy.db
	db, err := bolt.Open("D:\\ws_go\\learngo\\bolt_db\\fanal\\fanal.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// loop print bucket
	err = db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			fmt.Println("bucket:", string(name))
			return nil
		})
	})
	if err != nil {
		log.Fatal(err)
	}

	// 遍历 vulnerability bucket
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blob")) // vul db info
		//b := tx.Bucket([]byte("wolfi"))    // release version name
		if b == nil {
			return fmt.Errorf("bucket vulnerability not found")
		}

		// 游标遍历所有 key/value
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("Key = %s\n", k)

			// value 是 JSON，可以直接反序列化
			var data interface{}
			if err := json.Unmarshal(v, &data); err != nil {
				// 如果不是 JSON，直接打印原始字符串
				fmt.Printf("  Raw Value = %s\n", string(v))
			} else {
				// 格式化打印 JSON
				jsonBytes, _ := json.MarshalIndent(data, "  ", "  ")
				fmt.Printf("  JSON Value = %s\n", string(jsonBytes))
			}
			fmt.Println("------------------------------------------------")
		}
		return nil
	})

	//// Trivy 用 "vulnerability" bucket 存漏洞详情
	//err = db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte("vulnerability"))
	//	if b == nil {
	//		return fmt.Errorf("bucket not found")
	//	}
	//
	//	// 例子：读取 CVE-2023-12345
	//	vulnID := "CVE-2023-12345"
	//	val := b.Get([]byte(vulnID))
	//	if val != nil {
	//		fmt.Println(string(val)) // 这里是 JSON 格式
	//	} else {
	//		fmt.Println("not found:", vulnID)
	//	}
	//	return nil
	//})

	if err != nil {
		log.Fatal(err)
	}
}
