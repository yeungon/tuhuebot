package bbolt

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

// Define a bucket name.
var bucketName = []byte("users")

func CreateUserBucket() {
	// Ensure the bucket exists (only once).
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func UserAdd(key []byte, value []byte) {
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", []byte("users"))
		}
		// CREATE: Add a key-value pair.
		err = bucket.Put(key, value)
		if err != nil {
			return fmt.Errorf("put value: %s", err)
		}

		//fmt.Println("Value inserted")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
func UserRead(key []byte) []byte {
	var val []byte

	// READ: Retrieve a value.
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", "users")
		}

		// Get the value from the bucket.
		val = bucket.Get(key)
		if val == nil {
			return fmt.Errorf("key %q not found", key)
		}

		fmt.Printf("Read value: %s\n", val)
		fmt.Printf("Read value: %s\n", string(val))
		fmt.Printf("Read value: %v\n", val)
		return nil
	})

	if err != nil {
		log.Printf("Error reading from bucket: %v", err)
		return nil
	}

	return val
}

// PrintAllKeyValues prints all key-value pairs in a specified bucket.
func PrintAllKeyValues(bucketName string) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", bucketName)
		}

		// Use a cursor to iterate over all key-value pairs in the bucket.
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Printf("Key: %s, Value: %s\n", k, v)
		}

		return nil
	})

	if err != nil {
		log.Printf("Error reading from bucket %q: %v", bucketName, err)
	}
}

// CleanAllKeyValues deletes all key-value pairs in a specified bucket.
func CleanAllKeyValues(bucketName string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", bucketName)
		}

		// Use a cursor to iterate over all key-value pairs in the bucket.
		cursor := bucket.Cursor()
		for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {
			// Delete each key.
			if err := bucket.Delete(k); err != nil {
				return fmt.Errorf("failed to delete key %q: %v", k, err)
			}
		}

		fmt.Printf("All key-value pairs in bucket %q have been deleted.\n", bucketName)
		return nil
	})

	if err != nil {
		log.Printf("Error cleaning up bucket %q: %v", bucketName, err)
	}
}
