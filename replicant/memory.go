package replicant

import (
	"log"

	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func InitDB(path string) {
	var err error
	db, err = bbolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bbolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("traits"))
		tx.CreateBucketIfNotExists([]byte("knowledge"))
		return nil
	})
}

func Save(bucket, key, value string) {
	db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put([]byte(key), []byte(value))
	})
}

func Get(bucket, key string) string {
	var out []byte
	db.View(func(tx *bbolt.Tx) error {
		out = tx.Bucket([]byte(bucket)).Get([]byte(key))
		return nil
	})
	return string(out)
}

func CloseDB() {
	db.Close()
}
