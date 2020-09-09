package db

import "github.com/boltdb/bolt"

func getAllTasks() ([]Task, error) {
	tasks := make([]Task, 0)
	err :=  db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			//fmt.Printf("key=%d, value=%s\n", btoi(k), v)
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
