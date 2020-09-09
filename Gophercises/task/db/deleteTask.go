package db

import "github.com/boltdb/bolt"

func deleteTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b:= tx.Bucket([]byte(bucketName))
		return b.Delete(itob(id))
	})
}
