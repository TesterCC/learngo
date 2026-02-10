package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	bolt "go.etcd.io/bbolt"
)

func walkBucket(b *bolt.Bucket, prefix string) error {
	return b.ForEach(func(k, v []byte) error {
		if v == nil {
			// v == nil 说明这是子 bucket
			child := b.Bucket(k)
			if child == nil {
				return nil
			}
			fullName := filepath.Join(prefix, string(k))
			fmt.Println("Bucket:", fullName)
			return walkBucket(child, fullName)
		}

		// 普通 KV
		fmt.Printf("  Key: %s\n", k)

		var obj interface{}
		if err := json.Unmarshal(v, &obj); err == nil {
			data, _ := json.MarshalIndent(obj, "    ", "  ")
			fmt.Printf("  Value(JSON): %s\n", string(data))
		} else {
			fmt.Printf("  Value(Raw): %d bytes\n", len(v))
		}

		fmt.Println("  ---------------------------------------------")
		return nil
	})
}

func main() {
	dbPath := flag.String("db", "", "path to trivy bolt db file")
	flag.Parse()

	if *dbPath == "" {
		log.Fatal("please specify -db /path/to/trivy.db")
	}

	db, err := bolt.Open(*dbPath, 0600, &bolt.Options{
		ReadOnly: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			root := string(name)
			fmt.Println("Root Bucket:", root)
			return walkBucket(b, root)
		})
	})

	if err != nil {
		log.Fatal(err)
	}
}
