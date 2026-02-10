package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	bolt "go.etcd.io/bbolt"
)

func walkBucketSimple(b *bolt.Bucket, prefix string) error {
	return b.ForEach(func(k, v []byte) error {
		// v == nil 表示这是一个子 bucket
		if v != nil {
			return nil
		}

		child := b.Bucket(k)
		if child == nil {
			return nil
		}

		name := filepath.Join(prefix, string(k))
		fmt.Println(name)

		return walkBucketSimple(child, name)
	})
}

func main() {
	dbPath := flag.String("db", "", "path to trivy bolt db file")
	flag.Parse()

	if *dbPath == "" {
		log.Fatal("usage: bucketdump -db /path/to/trivy.db")
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
			fmt.Println(root)
			return walkBucketSimple(b, root)
		})
	})

	if err != nil {
		log.Fatal(err)
	}
}

/*
-db bolt_db/trivydb/db/trivy.db     # vulner

-db bolt_db/trivydb/fanal/trivy.db  # empty

-db bolt_db/trivydb/fanal/fanal.db
Root Bucket: artifact
Root Bucket: blob

-db bolt_db/trivydb/java-db/trivy-java.db  # 有数据，但很大，去mac上跑一下
#  2026/02/10 16:10:11 CreateFileMapping: Not enough memory resources are available to process this command.
# 在 Windows 上 mmap 这么大的文件非常容易失败, 推荐在Linux或docker中跑

*/
