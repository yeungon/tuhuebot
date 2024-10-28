package bbolt

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB
var err error

func Connect() *bolt.DB {
	db, err = bolt.Open("tuhuebot_kv.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	CreateUserBucket()
	return db
}

func BBolt() *bolt.DB {
	return db
}
