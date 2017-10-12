package main

import (
	"fmt"
	"log"

	bolt "github.com/coreos/bbolt"
)

func main() {
	viewDB("path/to/boltdb/file")
}

func viewDB(path string) {
	db, err := bolt.Open(path, 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	db.View(func(tx *bolt.Tx) error {
		tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			fmt.Printf("bucket: %s\n", string(name))
			viewBucketList(b)
			return nil
		})
		return nil
	})
}
func viewBucketList(b *bolt.Bucket) {
	b.ForEach(func(k, v []byte) error {
		if b.Bucket(k) != nil {
			fmt.Printf("bucket: %s\n", string(k))
			viewBucketList(b.Bucket(k))
		} else {
			fmt.Printf("key: %s, value: %s\n", string(k), string(v))
		}
		return nil
	})
}
