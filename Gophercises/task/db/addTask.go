package db

import "github.com/boltdb/bolt"

func addTask(task string) (int,error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		id64, _ := b.NextSequence()
		id = int(id64)
		return b.Put(itob(int(id64)), []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, err
}
